package api

import (
	"github.com/aleksrutins/litelytics/auth"
	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/aleksrutins/litelytics/ent/user"
	"github.com/aleksrutins/litelytics/ent/visit"
	"github.com/gofiber/fiber/v2"
)

type SiteData struct {
	ID   int `json:"id"`
	Site struct {
		Domain  string  `json:"domain"`
		Favicon *string `json:"favicon"`
	} `json:"site"`
	Visits []*ent.Visit `json:"visits"`
}

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
	siteInfo, err := dbutil.Client.Site.Query().
		Where(site.IDEQ(intId)).
		Where(site.HasUsersWith(user.IDEQ(currentUser.ID))).
		First(c.Context())

	if err != nil {
		return err
	}

	visits, err := dbutil.Client.Visit.Query().
		Where(visit.HasSiteWith(site.IDEQ(siteInfo.ID))).
		Order(ent.Asc(visit.FieldTimestamp)).
		All(c.Context())

	if err != nil {
		return err
	}

	c.JSON(SiteData{
		ID: intId,
		Site: struct {
			Domain  string  "json:\"domain\""
			Favicon *string "json:\"favicon\""
		}{Domain: siteInfo.Domain, Favicon: &siteInfo.Favicon},
		Visits: visits,
	})
	return nil
}
