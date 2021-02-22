package resources

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for any creed fetching commands
var creedsCommand = models.Command{
	Command: []string{
		"creeds",
		"apostles",
		"nicene325",
		"nicene",
		"chalcedon",
	},
	Params:         nil,
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Handles commands for the creeds and catechisms.
func router() {
	return
}
