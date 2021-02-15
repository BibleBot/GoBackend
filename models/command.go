package models

// Command is a basic command struct containing expected values.
type Command struct {
	Command        string
	Params         []string
	IsOwnerCommand bool

	Process func() error
}
