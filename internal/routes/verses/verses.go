package verses

// Parse message with parsing.go, then generate a reference here and process it.

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRouter registers routers related to verse processing.
func RegisterRouter(app *fiber.App) {
	app.Get("/api/verses/:input", verseHandler)
}

func verseHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("input"))
}
