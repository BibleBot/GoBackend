package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command stuct for the formatting command.
var formattingCommand = models.Command{
	Command: "formatting",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Handles all formatting-related commands - verse numbers, headings, and display styles.
func formatting() {
	return
}
