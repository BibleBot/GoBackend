package commands

// Just a place to point to the right command.
// Possibly create a Command struct to expect similar results? Might be easier said than done.

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
