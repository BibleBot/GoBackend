package models

// Command is an umbrella struct that can be used to hold info & process 1 or more commands
type Command struct {
	Command        []string
	Params         []string
	IsOwnerCommand bool

	Process func() error
}

// Basic help command (biblebot)
var Help = Command{
	Command:        []string{"biblebot"},
	Params:         nil,
	IsOwnerCommand: false,
	Process: func() error {
		return nil
	},
}
