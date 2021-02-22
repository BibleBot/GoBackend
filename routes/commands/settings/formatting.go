package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command stuct for the formatting command.
var formattingCommand = models.Command{
	Command: []string{
		"formatting",
	},
	Params: []string{
		"default",
		"embed",
		"blockquote",
		"code",
	},
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles all formatting-related commands - verse numbers, headings, and display styles.
func formatting() {
	return
}
