package api

import (
	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/aleksrutins/litelytics/ent/user"
	"github.com/gofiber/fiber/v2"
)

func sites(c *fiber.Ctx) error {
	currentUser := auth.GetUser(c)
	if currentUser == nil {
		c.JSON(fiber.Map{"Error": "Not authenticated"})
		return fiber.ErrUnauthorized
	}
	sites, err := dbutil.Client.Site.Query().Where(site.HasUsersWith(user.IDEQ(currentUser.ID))).All(c.Context())
	if err != nil {
		return err
	}
	c.JSON(sites)
	return nil
}
