package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// VerseCommandRouter is a basic struct with functions to handle verse-related commands.
type VerseCommandRouter struct {
	Commands []models.Command
}

var (
	// verseInstance is the singleton router used to process its respective commands
	verseInstance *VerseCommandRouter
	verseOnce     sync.Once

	searchCommand = models.Command{
		Command: "search",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
	randomCommand = models.Command{
		Command: "random",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
	truerandomCommand = models.Command{
		Command: "truerandom",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
)

// NewVerseCommandRouter creates a VerseCommandRouter if one does not already exist
func NewVerseCommandRouter() *VerseCommandRouter {
	verseOnce.Do(func() {
		verseInstance = &VerseCommandRouter{
			Commands: []models.Command{
				searchCommand,
				randomCommand,
				truerandomCommand,
			},
		}
	})

	return verseInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *VerseCommandRouter) Process(params []string, ctx *models.Context) (*models.CommandResponse, error) {
	cm, ok := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:], ctx)
	}
	return nil, nil // Implement return error
}
