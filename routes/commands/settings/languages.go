package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// LanguageCommandRouter is a basic struct with functions to handle language-related commands.
type LanguageCommandRouter struct {
	Commands []models.Command
}

var (
	// languageInstance is the singleton router used to process its respective commands
	languageInstance *LanguageCommandRouter
	languageOnce     sync.Once

	languagesCommand = models.Command{
		Command: "language",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewLanguageCommandRouter creates a LanguageCommandRouter if one does not already exist
func NewLanguageCommandRouter() *LanguageCommandRouter {
	languageOnce.Do(func() {
		languageInstance = &LanguageCommandRouter{
			Commands: []models.Command{languagesCommand},
		}
	})

	return languageInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *LanguageCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}
