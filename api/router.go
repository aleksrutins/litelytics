package api

import "github.com/gofiber/fiber/v2"

var Routes = fiber.New()

func init() {
	Routes.
		Get("/sites", sites)
}
