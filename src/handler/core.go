package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Context struct{}

func Init(app *fiber.App) {
	ctx := &Context{}

	app.Get("/", ctx.Health)
}
