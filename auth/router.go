package auth

import "github.com/gofiber/fiber/v2"

var Routes = fiber.New()

func init() {
	Routes.
		Get("/login", loginPage).
		Get("/logout", logOut).
		Post("/login", authenticate).
		Post("/create-account", createUser)
}
