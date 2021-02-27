package information

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type HelpCommandRouter struct {
	Commands []models.Command
	Process  func([]string) error
}

var (
	helpOnce     sync.Once
	helpInstance *HelpCommandRouter

	helpCommand = models.Command{
		Command: "biblebot",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewHelpCommandRouter creates a HelpCommandRouter if one does not exist
func NewHelpCommandRouter() *HelpCommandRouter {
	helpOnce.Do(func() {
		helpInstance = &HelpCommandRouter{
			Commands: []models.Command{helpCommand},
		}
	})

	return helpInstance
}

// RouterProcess checks which command process to run given the inputed command & parameters
func (hcr *HelpCommandRouter) RouterProcess(params []string) error {
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
