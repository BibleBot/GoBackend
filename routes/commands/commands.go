package commands

// Just a place to point to the right command.
// Possibly create a Command struct to expect similar results? Might be easier said than done.

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
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
	commandList := make(map[string][]string)

	// Check if command is valid
	// If we're testing, the working directory is tests/, so paths need to be adjusted for that.
	dir := "./"
	if _, err := os.Stat(dir + "data/names/command_list.json"); os.IsNotExist(err) {
		dir = "./../"
	}

	file, err := ioutil.ReadFile(dir + "data/names/command_list.json")
	if err != nil {
		panic(logger.LogWithError("commands", "could not read command_list.json", nil))
	}
	json.Unmarshal(file, &commandList)

	return slices.StringInSlice(cm, commandList["commands"]) || slices.StringInSlice(cm, commandList["owner_commands"])
}
