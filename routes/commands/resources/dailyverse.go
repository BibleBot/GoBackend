package resources

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for the dailyverse command.
var dailyverseCommand = models.Command{
	Command: "dailyverse",
	Process: func() error {
		return nil // To implement
	},
}

// Handles the daily verse command and anything related to it.
func dailyVerse() {
	return
}
