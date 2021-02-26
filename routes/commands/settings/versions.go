package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for changing versions.
var versionCommand = models.Command{
	Command: "version",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Handles all version-related commands.
func versions() {
	return
}
