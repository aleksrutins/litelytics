package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func logOut(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{Name: "userId", Expires: time.Now().Add(-24 * time.Hour)})
	c.Cookie(&fiber.Cookie{Name: "userEmail", Expires: time.Now().Add(-24 * time.Hour)})
	c.RedirectBack("/")
	return nil
}
