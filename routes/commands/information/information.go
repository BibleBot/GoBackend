package information

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// HelpCommandRouter is a basic struct with functions to handle help-related commands.
type HelpCommandRouter struct {
	Commands []models.Command
}

var (
	// helpInstance is the singleton router used to process its respective commands
	helpInstance *HelpCommandRouter
	helpOnce     sync.Once

	helpCommand = models.Command{
		Command: "biblebot",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
)

// NewHelpCommandRouter creates a HelpCommandRouter if one does not already exist
func NewHelpCommandRouter() *HelpCommandRouter {
	helpOnce.Do(func() {
		helpInstance = &HelpCommandRouter{
			Commands: []models.Command{helpCommand},
		}
	})

	return helpInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *HelpCommandRouter) Process(params []string, ctx *models.Context) (*models.CommandResponse, error) {
	cm, ok := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:], ctx)
	}
	return nil, nil // Implement return error
}
