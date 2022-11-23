package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/aleksrutins/litelytics/auth"
)

func main() {
	templates := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: templates,
	})

	app.Static("/static", "./static")

	app.Mount("/auth", auth.Routes)
	port := os.Getenv("PORT")
	if port == "" { port = "3000" }
	app.Listen("0.0.0.0:" + port)
}
