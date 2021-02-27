package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type formattingCommandRouter struct {
	Commands []models.Command
}

var (
	// FormattingInstance is the singleton router used to process its respective commands
	FormattingInstance *formattingCommandRouter
	formattingOnce     sync.Once

	formattingCommand = models.Command{
		Command: "formatting",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewFormattingCommandRouter creates a formattingCommandRouter if one does not already exist
func NewFormattingCommandRouter() *formattingCommandRouter {
	formattingOnce.Do(func() {
		FormattingInstance = &formattingCommandRouter{
			Commands: []models.Command{formattingCommand},
		}
	})

	return FormattingInstance
}

// Process checks which command process to run given the inputed command & parameters
func (fcr *formattingCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(fcr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}

// Handles all formatting-related commands - verse numbers, headings, and display styles.
func formatting() {
	return
}
