package settings

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// VersionCommandRouter is a basic struct with functions to handle version-related commands.
type VersionCommandRouter struct {
	Commands []models.Command
}

var (
	// versionInstance is the singleton router used to process its respective commands
	versionInstance *VersionCommandRouter
	versionOnce     sync.Once

	versionsCommand = models.Command{
		Command: "version",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewVersionCommandRouter creates a VersionCommandRouter if one does not already exist
func NewVersionCommandRouter() *VersionCommandRouter {
	versionOnce.Do(func() {
		versionInstance = &VersionCommandRouter{
			Commands: []models.Command{versionsCommand},
		}
	})

	return versionInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *VersionCommandRouter) Process(params []string) error {
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
