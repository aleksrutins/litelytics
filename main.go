package main

import (
	"log"
	"os"

	"github.com/aleksrutins/litelytics/api"
	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html"
	"github.com/profclems/go-dotenv"
)

func main() {
	err := dotenv.LoadConfig()

	if err != nil {
		log.Printf(".env could not be loaded: %v\n", err)
	}

	dbutil.Connect()
	defer dbutil.Client.Close()

	templates := html.New("./templates", ".html")

	templates.AddFunc("headMeta", func(title string) map[string]interface{} {
		return fiber.Map{"Title": title}
	})

	app := fiber.New(fiber.Config{
		Views: templates,
	})

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: dotenv.GetString("SECRET_KEY"),
	}))

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		if auth.GetUser(c) == nil {
			c.Redirect("/auth/login")
		}
		return c.Render("index", util.CreateContext(c))
	})
	app.Get("/sites/:id", func(c *fiber.Ctx) error {
		if auth.GetUser(c) == nil {
			c.Redirect("/auth/login")
		}
		return c.Render("site_info", util.CreateContext(c))
	})

	app.Mount("/auth", auth.Routes)
	app.Mount("/api", api.Routes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen("0.0.0.0:" + port)
}
