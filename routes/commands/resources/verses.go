package resources

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for verse lookup commands.
var verseFetchCommand = models.Command{
	Command: []string{
		"search",
		"random",
		"truerandom",
	},
	Params:         nil, // Valid reference if applicable
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles verse lookups and random/truerandom verses.
func verses() {
	return
}
