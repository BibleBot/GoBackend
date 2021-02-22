package models

// Command is an umbrella struct that can be used to hold info & process 1 or more commands
type Command struct {
	Command        []string
	Params         []string
	IsOwnerCommand bool

	Process func() error
}

/* 2 categories homeless at the moment
// Commands for fetching miscellaneous info
var MiscFetch = Command{
	Command: []string{
		"invite",
		"stats",
		"misc",
		"ping",
	},
	Params:         nil,
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Commands only executable by owner(s)
var Owner = Command{
	Command: []string{
		"echo",
		"eval",
		"leave",
	},
	Params:         nil, // Statements
	IsOwnerCommand: true,
	Process: func() error {
		return nil // To implement
	},
}*/
