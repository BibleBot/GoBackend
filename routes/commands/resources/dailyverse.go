package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// DailyVerseCommandRouter is a basic struct with functions to handle daily verse-related commands.
type DailyVerseCommandRouter struct {
	Commands []models.Command
}

var (
	// dailyVerseInstance is the singleton router used to process its respective commands
	dailyVerseInstance *DailyVerseCommandRouter
	dailyVerseOnce     sync.Once

	dailyVerseCommand = models.Command{
		Command: "dailyverse",
		Process: func(params []string, ctx *models.Context) (*models.CommandResponse, error) {
			return nil, nil // To implement
		},
	}
)

// NewDailyVerseCommandRouter creates a DailyVerseCommandRouter if one does not already exist
func NewDailyVerseCommandRouter() *DailyVerseCommandRouter {
	dailyVerseOnce.Do(func() {
		dailyVerseInstance = &DailyVerseCommandRouter{
			Commands: []models.Command{dailyVerseCommand},
		}
	})

	return dailyVerseInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *DailyVerseCommandRouter) Process(params []string, ctx *models.Context) (*models.CommandResponse, error) {
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
