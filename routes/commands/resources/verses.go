package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type versesCommandRouter struct {
	Commands []models.Command
}

var (
	// VersesInstance is the singleton router used to process its respective commands
	VersesInstance *versesCommandRouter
	versesOnce     sync.Once

	searchCommand = models.Command{
		Command: "search",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	randomCommand = models.Command{
		Command: "random",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
	truerandomCommand = models.Command{
		Command: "truerandom",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewVersesCommandRouter creates a versesCommandRouter if one does not already exist
func NewVersesCommandRouter() *versesCommandRouter {
	versesOnce.Do(func() {
		VersesInstance = &versesCommandRouter{
			Commands: []models.Command{
				searchCommand,
				randomCommand,
				truerandomCommand,
			},
		}
	})

	return VersesInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *versesCommandRouter) Process(params []string) error {
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

// Handles verse lookups and random/truerandom verses.
func verses() {
	return
}
