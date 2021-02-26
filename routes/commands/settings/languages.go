package settings

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for changing languages.
var languageCommand = models.Command{
	Command: "language",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Handles all language-related commands. (Should be quite similar to versions.go as they basically do the same thing, give or take some)
func languages() {
	return
}
