package commands

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRouter registers routers related to command processing.
func RegisterRouter(app *fiber.App) {
	app.Get("/api/commands/:input/:args?", commandHandler)
}

func commandHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("input"))
}
