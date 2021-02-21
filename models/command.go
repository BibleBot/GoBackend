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
	Command: []string{
		"biblebot",
	},
	Params:         nil,
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Commands for changing formatting
var Formatting = Command{
	Command: []string{
		"formatting",
	},
	Params: []string{
		"default",
		"embed",
		"blockquote",
		"code",
	},
	IsOwnerCommand: false,
	Process: func() error {
		return nil // To implement
	},
}

// Commands for changing versions/language
var Translations = Command{
	Command: []string{
		"version",
		"language",
	},
	Params:         nil, // Valid version or language using go:embed? or initialize or smth
	IsOwnerCommand: false,
	Process: func() error {
		return nil
	},
}
