package api

import (
	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/aleksrutins/litelytics/ent/user"
	"github.com/gofiber/fiber/v2"
)

func siteData(c *fiber.Ctx) error {
	currentUser := auth.GetUser(c)
	if currentUser == nil {
		c.JSON(fiber.Map{"Error": "Not authenticated"})
		return fiber.ErrUnauthorized
	}

	intId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	site, err := dbutil.Client.Site.Query().
		Where(site.IDEQ(intId)).
		Where(site.HasUsersWith(user.IDEQ(currentUser.ID))).First(c.Context())
	if err != nil {
		return err
	}
	c.JSON(site)
	return nil
}
