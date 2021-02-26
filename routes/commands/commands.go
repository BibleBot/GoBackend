package commands

// Just a place to point to the right command.
// Possibly create a Command struct to expect similar results? Might be easier said than done.

// Homeless Commands: {"invite", "stats", "misc", "ping", "echo", "eval", "leave"}

import (
	_ "embed" // for go:embeds
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

var (
	//go:embed data/command_list.json
	commandListFile []byte

	commandList = make(map[string][]string)
)

// RegisterRouter registers routers related to command processing.
func RegisterRouter(app *fiber.App) {
	app.Get("/api/commands/:input/:args?", commandHandler)
}

func commandHandler(c *fiber.Ctx) error {
	// Check prefix before creating command struct and calling Process()

	return c.SendString(c.Params("input"))
}

func isValidCommand(cm string) bool {
	if len(commandList) == 0 {
		json.Unmarshal(commandListFile, &commandList)
	}

	return slices.StringInSlice(cm, commandList["commands"]) || slices.StringInSlice(cm, commandList["owner_commands"])
}
