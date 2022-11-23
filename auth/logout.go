package auth

import "github.com/gofiber/fiber/v2"

func logOut(c *fiber.Ctx) error {
	c.ClearCookie("userId", "userEmail")
	return nil
}
