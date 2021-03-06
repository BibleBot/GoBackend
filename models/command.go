package models

// Command is a struct that holds a type of command & the relevant process for it
type Command struct {
	Command string

	Process func([]string, *Context) *CommandResponse
}
