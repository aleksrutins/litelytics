package util

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// https://github.com/gofiber/fiber/issues/299

func WrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
		return nil
	}
}
