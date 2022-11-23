package auth

import "github.com/gofiber/fiber/v2"

var Routes = fiber.New()

func init() {
	Routes.Get("/login", loginPage)
}
