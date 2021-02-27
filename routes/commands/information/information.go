package information

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type helpCommandRouter struct {
	Commands []models.Command
}

var (
	// HelpInstance is the singleton router used to process its respective commands
	HelpInstance *helpCommandRouter
	helpOnce     sync.Once

	helpCommand = models.Command{
		Command: "biblebot",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewHelpCommandRouter creates a helpCommandRouter if one does not already exist
func NewHelpCommandRouter() *helpCommandRouter {
	helpOnce.Do(func() {
		HelpInstance = &helpCommandRouter{
			Commands: []models.Command{helpCommand},
		}
	})

	return HelpInstance
}

// Process checks which command process to run given the inputed command & parameters
func (hcr *helpCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(hcr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}

// Handles help commands.
func router() {
	return
}
