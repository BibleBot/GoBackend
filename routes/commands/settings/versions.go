package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for changing versions.
var versionCommand = models.Command{
	Command: []string{
		"version",
	},
	Params:         nil, // Valid version or language using go:embed? or initialize or smth
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles all version-related commands.
func versions() {
	return
}
