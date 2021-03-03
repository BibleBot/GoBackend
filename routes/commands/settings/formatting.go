package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// FormatCommandRouter is a basic struct with functions to handle format-related commands.
type FormatCommandRouter struct {
	Commands []models.Command
}

var (
	// formattingInstance is the singleton router used to process its respective commands
	formattingInstance *FormatCommandRouter
	formattingOnce     sync.Once

	formattingCommand = models.Command{
		Command: "formatting",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
)

// NewFormattingCommandRouter creates a FormatCommandRouter if one does not already exist
func NewFormattingCommandRouter() *FormatCommandRouter {
	formattingOnce.Do(func() {
		formattingInstance = &FormatCommandRouter{
			Commands: []models.Command{formattingCommand},
		}
	})

	return formattingInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *FormatCommandRouter) Process(params []string, ctx *models.Context) (*models.CommandResponse, error) {
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
