package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// ResourceCommandRouter is a basic struct with functions to handle resource-related commands.
type ResourceCommandRouter struct {
	Commands []models.Command
}

var (
	// resourceInstance is the singleton router used to process its respective commands
	resourceInstance *ResourceCommandRouter
	resourceOnce     sync.Once
)

// NewResourceCommandRouter creates a ResourceCommandRouter if one does not already exist
func NewResourceCommandRouter() *ResourceCommandRouter {
	resourceOnce.Do(func() {
		resourceInstance = &ResourceCommandRouter{
			Commands: []models.Command{
				creedsCommand,
				apostlesCommand,
				nicene325Command,
				niceneCommand,
				chalcedonCommand,
			},
		}
	})

	return resourceInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *ResourceCommandRouter) Process(params []string, ctx *models.Context) *models.CommandResponse {
	cm, ok := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:], ctx)
	}
	return nil
}

var creedsCommand = models.Command{
	Command: "creeds",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}

var apostlesCommand = models.Command{
	Command: "apostles",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}

var nicene325Command = models.Command{
	Command: "nicene325",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}

var niceneCommand = models.Command{
	Command: "nicene",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}

var chalcedonCommand = models.Command{
	Command: "chalcedon",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		return nil
	},
}
