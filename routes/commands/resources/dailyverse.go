package resources

import (
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

type dailyVerseCommandRouter struct {
	Commands []models.Command
}

var (
	// DailyVerseInstance is the singleton router used to process its respective commands
	DailyVerseInstance *dailyVerseCommandRouter
	dailyVerseOnce     sync.Once

	dailyVerseCommand = models.Command{
		Command: "dailyverse",
		Process: func(params []string) error {
			return nil // To implement
		},
	}
)

// NewDailyVerseCommandRouter creates a dailyVerseCommandRouter if one does not already exist
func NewDailyVerseCommandRouter() *dailyVerseCommandRouter {
	dailyVerseOnce.Do(func() {
		DailyVerseInstance = &dailyVerseCommandRouter{
			Commands: []models.Command{dailyVerseCommand},
		}
	})

	return DailyVerseInstance
}

// Process checks which command process to run given the inputed command & parameters
func (dvcr *dailyVerseCommandRouter) Process(params []string) error {
	cm, ok := slices.FilterInterface(dvcr.Commands, func(cm interface{}) bool {
		cmd, ok := cm.(models.Command)
		return (params[0] == cmd.Command) && ok
	}).(models.Command)

	if ok {
		// Strip first element of slice (is the command itself)
		return cm.Process(params[1:])
	}
	return nil // Implement return error
}

// Handles the daily verse command and anything related to it.
func dailyVerse() {
	return
}
