package verses

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App) {
	app.Get("/api/verses/:input", verseHandler)
}

func verseHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("input"));
}
