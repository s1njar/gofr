package handler

import "github.com/gofiber/fiber/v2"

func (ctx *Context) Health(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
