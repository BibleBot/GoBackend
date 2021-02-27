package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type resourcesCommandRouter struct {
	Commands []models.Command
}

var (
	// ResourcesInstance is the singleton router used to process its respective commands
	ResourcesInstance *resourcesCommandRouter
	resourcesOnce     sync.Once

	creedsCommand = models.Command{
		Command: "creeds",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	apostlesCommand = models.Command{
		Command: "apostles",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	nicene325Command = models.Command{
		Command: "nicene325",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	niceneCommand = models.Command{
		Command: "nicene",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	chalcedonCommand = models.Command{
		Command: "chalcedon",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewResourcesCommandRouter creates a resourcesCommandRouter if one does not already exist
func NewResourcesCommandRouter() *resourcesCommandRouter {
	resourcesOnce.Do(func() {
		ResourcesInstance = &resourcesCommandRouter{
			Commands: []models.Command{
				creedsCommand,
				apostlesCommand,
				nicene325Command,
				niceneCommand,
				chalcedonCommand,
			},
		}
	})

	return ResourcesInstance
}

// Process checks which command process to run given the inputed command & parameters
func (rcr *resourcesCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(rcr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}

// Handles commands for the creeds and catechisms.
func router() {
	return
}
