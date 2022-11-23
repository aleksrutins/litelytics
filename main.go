package main

import (
	"log"
	"os"

	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
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

	app := fiber.New(fiber.Config{
		Views: templates,
	})

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: dotenv.GetString("SECRET_KEY"),
	}))

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		if !auth.IsAuthenticated(c) {
			c.Redirect("/auth/login")
		}
		c.Send([]byte("Welcome"))
		return nil
	})

	app.Mount("/auth", auth.Routes)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen("0.0.0.0:" + port)
}
