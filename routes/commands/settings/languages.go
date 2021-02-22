package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for changing languages.
var languageCommand = models.Command{
	Command: []string{
		"language",
	},
	Params:         nil, // Valid version or language using go:embed? or initialize or smth
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles all language-related commands. (Should be quite similar to versions.go as they basically do the same thing, give or take some)
func languages() {
	return
}
