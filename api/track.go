package api

import (
	"net/url"
	"time"

	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/gofiber/fiber/v2"
)

type trackRequest struct {
	Referrer string `json:"referrer"`
}

func track(c *fiber.Ctx) error {
	var req trackRequest
	c.BodyParser(req)
	requestURL, err := url.Parse(string(c.Context().Referer()))
	if err != nil {
		return err
	}

	site, err := dbutil.Client.Site.Query().Where(site.DomainEQ(requestURL.Host)).First(c.Context())
	if err != nil {
		return nil
	}

	_, err = dbutil.Client.Visit.
		Create().
		SetPath(requestURL.Path).
		SetReferrer(req.Referrer).
		SetIP(c.IP()).
		SetTimestamp(time.Now()).
		SetSite(site).
		Save(c.Context())

	return err
}
