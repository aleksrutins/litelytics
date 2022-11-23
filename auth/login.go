package auth

import "github.com/gofiber/fiber/v2"

func loginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"Error": nil})
}
