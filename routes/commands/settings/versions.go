package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type versionsCommandRouter struct {
	Commands []models.Command
}

var (
	// VersionsInstance is the singleton router used to process its respective commands
	VersionsInstance *versionsCommandRouter
	versionsOnce     sync.Once

	versionsCommand = models.Command{
		Command: "version",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewVersionsCommandRouter creates a versionsCommandRouter if one does not already exist
func NewVersionsCommandRouter() *versionsCommandRouter {
	versionsOnce.Do(func() {
		VersionsInstance = &versionsCommandRouter{
			Commands: []models.Command{versionsCommand},
		}
	})

	return VersionsInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *versionsCommandRouter) Process(params []string) error {
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

// Handles all version-related commands.
func versions() {
	return
}
