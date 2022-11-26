package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/aleksrutins/litelytics/api"
	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html"
	"github.com/profclems/go-dotenv"
	vueglue "github.com/torenware/vite-go"
)

//go:embed "frontend"
var frontend embed.FS

var vueGlue *vueglue.VueGlue

func main() {
	err := dotenv.LoadConfig()
	if err != nil {
		log.Printf(".env could not be loaded: %v\n", err)
	}

	var viteConfig *vueglue.ViteConfig

	if util.IsProduction() {
		fmt.Println("Running in production.")
		frontendFS, err := fs.Sub(frontend, "frontend")
		if err != nil {
			log.Fatalf("Error opening embedded frontend FS: %v", err)
		}
		viteConfig = &vueglue.ViteConfig{
			Environment: "production",
			AssetsPath:  "dist",
			EntryPoint:  "src/main.ts",
			Platform:    "vue",
			FS:          frontendFS,
		}
	} else {
		fmt.Println("Running in development.")
		viteConfig = &vueglue.ViteConfig{
			Environment: "development",
			AssetsPath:  "frontend",
			EntryPoint:  "src/main.ts",
			Platform:    "vue",
			FS:          os.DirFS("frontend"),
		}
	}

	vueGlue, err = vueglue.NewVueGlue(viteConfig)

	if err != nil {
		log.Fatalf("error initializing Vite: %v", err)
	}

	dbutil.Connect()
	defer dbutil.Client.Close()

	templates := html.New("./templates", ".html")

	templates.AddFunc("headMeta", func(title string) map[string]interface{} {
		return fiber.Map{"Title": title}
	})

	app := fiber.New(fiber.Config{
		Views:                   templates,
		EnableTrustedProxyCheck: util.IsProduction(),
	})

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    dotenv.GetString("SECRET_KEY"),
		Except: []string{"userEmail"},
	}))

	fileServer, err := vueGlue.FileServer()

	if err != nil {
		log.Fatalf("error creating static file server: %v", err)
	}

	app.Mount("/auth", auth.Routes)
	app.Mount("/api", api.Routes)

	app.Use("/api/track", cors.New(cors.Config{AllowMethods: "POST"}))
	app.Use("/simpleclient.js", cors.New(cors.Config{AllowMethods: "GET"}))

	app.Use(viteConfig.URLPrefix, util.WrapHandler(fileServer.ServeHTTP))
	if !util.IsProduction() {
		app.Static("/", "frontend/public")
	}

	app.Get("/*", func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/auth") || strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return c.Render("embedVue", vueGlue)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen("0.0.0.0:" + port)
}
