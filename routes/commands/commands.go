package commands

// Just a place to point to the right command.
// Possibly create a Command struct to expect similar results? Might be easier said than done.

// Homeless Commands: {"invite", "stats", "misc", "ping", "echo", "eval", "leave"}

import (
	_ "embed" // for go:embeds
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/routes/commands/settings"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/converters"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

var (
	//go:embed data/command_list.json
	commandListFile []byte

	commandList = make(map[string][]string)

	config *models.Config

	vcr = settings.NewVersionCommandRouter()
)

// RegisterRouter registers routers related to command processing.
func RegisterRouter(app *fiber.App, cfg *models.Config) {
	config = cfg

	app.Get("/api/commands/process", commandHandler)
}

func commandHandler(c *fiber.Ctx) error {
	// Check prefix before creating command struct and calling Process()
	ctx, err := converters.InputToContext(c.Body(), config)
	if err != nil {
		if err == fmt.Errorf("unauth") {
			c.SendStatus(401)
		} else {
			c.SendStatus(400)
		}

		return err
	}

	res := vcr.Process(strings.Split(ctx.Body, " ")[1:], ctx)

	return c.JSON(res)
}

func isValidCommand(cm string) bool {
	if len(commandList) == 0 {
		json.Unmarshal(commandListFile, &commandList)
	}

	return slices.StringInSlice(cm, commandList["commands"]) || slices.StringInSlice(cm, commandList["owner_commands"])
}
