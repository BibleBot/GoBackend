package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type languagesCommandRouter struct {
	Commands []models.Command
}

var (
	// LanguagesInstance is the singleton router used to process its respective commands
	LanguagesInstance *languagesCommandRouter
	languagesOnce     sync.Once

	languagesCommand = models.Command{
		Command: "language",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewLanguagesCommandRouter creates a LanguagesCommandRouter if one does not already exist
func NewLanguagesCommandRouter() *languagesCommandRouter {
	languagesOnce.Do(func() {
		LanguagesInstance = &languagesCommandRouter{
			Commands: []models.Command{languagesCommand},
		}
	})

	return LanguagesInstance
}

// Process checks which command process to run given the inputed command & parameters
func (hcr *languagesCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(hcr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}

// Handles all language-related commands. (Should be quite similar to versions.go as they basically do the same thing, give or take some)
func languages() {
	return
}
