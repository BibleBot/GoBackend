package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// FormatCommandRouter is a basic struct with functions to handle format-related commands.
type FormatCommandRouter struct {
	DefaultCommand models.Command
	Commands       []models.Command
}

var (
	// formattingInstance is the singleton router used to process its respective commands
	formattingInstance *FormatCommandRouter
	formattingOnce     sync.Once
)

// NewFormattingCommandRouter creates a FormatCommandRouter if one does not already exist
func NewFormattingCommandRouter() *FormatCommandRouter {
	formattingOnce.Do(func() {
		formattingInstance = &FormatCommandRouter{
			DefaultCommand: fmtDefault,
			Commands:       []models.Command{},
		}
	})

	return formattingInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *FormatCommandRouter) Process(params []string, ctx *models.Context) *models.CommandResponse {
	cmMatches := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		if len(params) > 0 {
			cmd, ok := cm.(models.Command)
			return (params[0] == cmd.Command) && ok
		}

		return false
	}).([]models.Command)

	var cm models.Command

	if len(cmMatches) == 0 {
		cm = cr.DefaultCommand
	} else {
		cm = cmMatches[0]
	}

	if len(params) == 0 {
		return cm.Process([]string{}, ctx)
	}

	return cm.Process(params[1:], ctx)
}

// Get user preferences (headings, verse numbers, display style) and a list of formatting commands
var fmtDefault = models.Command{
	Command: "formatting",
	Process: func(_ []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}

// Get guild preferences (prefix, brackets)

// Set guild brackets (<>, [], {}, ()) <- allow more than one?

// Set user headings (true, false)

// Set user verse numbers (true, false)

// Set user display style (default, embed, blockquote, code)
