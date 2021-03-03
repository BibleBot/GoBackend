package settings

import (
	"errors"
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// VersionCommandRouter is a basic struct with functions to handle version-related commands.
type VersionCommandRouter struct {
	DefaultCommand models.Command
	Commands       []models.Command
}

var (
	// versionInstance is the singleton router used to process its respective commands
	versionInstance *VersionCommandRouter
	versionOnce     sync.Once
)

// NewVersionCommandRouter creates a VersionCommandRouter if one does not already exist
func NewVersionCommandRouter() *VersionCommandRouter {
	versionOnce.Do(func() {
		versionInstance = &VersionCommandRouter{
			DefaultCommand: cmVersion,
			Commands:       []models.Command{},
		}
	})

	return versionInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *VersionCommandRouter) Process(params []string, ctx *models.Context) (*models.CommandResponse, error) {
	cm, ok := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:], ctx)
	}

	return nil, errors.New("no correlating version command") // Implement return error
}

var cmVersion = models.Command{
	Command: "version",
	Process: func(_ []string, ctx *models.Context) (*models.CommandResponse, error) {
		return nil, nil // to be implemented
	},
}
