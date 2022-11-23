package util

import (
	"github.com/aleksrutins/litelytics/auth"
	"github.com/gofiber/fiber/v2"
)

type NavItem struct {
	URL   string
	Label string
}

var DefaultNav = []NavItem{
	{URL: "/", Label: "Sites"},
}

type navItemInfo struct {
	Active bool
	NavItem
}
type DashboardContext struct {
	Me    *auth.UserInfo
	Nav   []navItemInfo
	Error *string
}

func CreateContext(c *fiber.Ctx) DashboardContext {
	var ctx DashboardContext
	ctx.Me = auth.GetUser(c)
	for _, item := range DefaultNav {
		ctx.Nav = append(ctx.Nav, navItemInfo{
			NavItem: item,
			Active:  c.Route().Path == item.URL,
		})
	}
	return ctx
}
