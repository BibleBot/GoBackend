package resources

import "internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

// Command struct for the creeds command.
var creedsCommand = models.Command{
	Command: "creeds",
	Process: func() error {
		return nil // To implement
	},
}

// Command struct for the apostles command.
var apostlesCommand = models.Command{
	Command: "apostles",
	Process: func() error {
		return nil // To implement
	},
}

// Command struct for the nicene325 command.
var nicene325Command = models.Command{
	Command: "nicene325",
	Process: func() error {
		return nil // To implement
	},
}

// Command struct for the nicene command.
var niceneCommand = models.Command{
	Command: "nicene",
	Process: func() error {
		return nil // To implement
	},
}

// Command struct for the chalcedon command.
var chalcedonCommand = models.Command{
	Command: "chalcedon",
	Process: func() error {
		return nil // To implement
	},
}

// Handles commands for the creeds and catechisms.
func router() {
	return
}
