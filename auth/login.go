package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func loginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"Error": nil})
}

func authenticate(c *fiber.Ctx) error {
	var req Credentials
	err := c.BodyParser(&req)
	if err != nil {
		log.Printf("Error parsing login credentials: %v\n", err)
	}

	user, success := checkCredentials(c.Context(), req.Email, req.Password)

	if !success {
		return c.Render("login", fiber.Map{"Error": "Invalid credentials"})
	}

	authenticateRequest(c, user)

	return c.Redirect("/")
}
