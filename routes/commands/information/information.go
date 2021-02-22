package information

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for help commands.
var helpCommand = models.Command{
	Command: []string{
		"biblebot",
	},
	Params:         nil,
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles help commands.
func router() {
	return
}
