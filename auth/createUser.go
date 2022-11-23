package auth

import (
	"crypto/sha256"

	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent/user"
	"github.com/gofiber/fiber/v2"
)

func createUser(c *fiber.Ctx) error {
	var req Credentials
	c.BodyParser(&req)
	if dbutil.Client.User.Query().Where(user.EmailEQ(req.Email)).ExistX(c.Context()) {
		c.Render("login", fiber.Map{"Error": "User already exists"})
		return nil
	}
	hash := sha256.Sum256([]byte(req.Password))
	user, err := dbutil.Client.User.Create().SetEmail(req.Email).SetPassword(hash[:]).Save(c.Context())
	if err != nil {
		c.Render("login", fiber.Map{"Error": "Error creating user"})
		return err
	}
	authenticateRequest(c, user)
	c.Redirect("/")
	return nil
}
