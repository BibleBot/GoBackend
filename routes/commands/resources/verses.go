package resources

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for the search command.
var searchCommand = models.Command{
	Command: "search",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Command struct for the random command.
var randomCommand = models.Command{
	Command: "random",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Command struct for the truerandom command.
var truerandomCommand = models.Command{
	Command: "truerandom",
	Process: func(params []string) error {
		return nil // To implement
	},
}

// Handles verse lookups and random/truerandom verses.
func verses() {
	return
}
