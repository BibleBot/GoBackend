package settings

import (
	"fmt"
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
			DefaultCommand: cmDefault,
			Commands:       []models.Command{cmSet},
		}
	})

	return versionInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *VersionCommandRouter) Process(params []string, ctx *models.Context) *models.CommandResponse {
	cmMatches := slices.FilterInterface(cr.Commands, func(cm interface{}) bool {
		if len(params) > 0 {
			cmd, ok := cm.(models.Command)
			return (params[0] == cmd.Command) && ok
		} else {
			return false
		}
	}).([]models.Command)

	var cm models.Command

	if len(cmMatches) == 0 {
		cm = cr.DefaultCommand
	} else {
		cm = cmMatches[0]
	}

	if len(params) > 0 {
		return cm.Process(params[1:], ctx)
	} else {
		return cm.Process([]string{}, ctx)
	}
}

var cmDefault = models.Command{
	Command: "version",
	Process: func(_ []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		var userVersion models.Version
		var guildVersion models.Version

		ctx.DB.Where(&models.Version{Abbreviation: ctx.Prefs.Version}).First(&userVersion)
		ctx.DB.Where(&models.Version{Abbreviation: ctx.GuildPrefs.Version}).First(&guildVersion)

		if ctx.Prefs.Version == "" {
			userVersion.Name = "undefined"
		}

		if ctx.GuildPrefs.Version == "" {
			guildVersion.Name = "undefined"
		}

		content := fmt.Sprintf("You are using **%s**.\nThis server is using **%s**.", userVersion.Name, guildVersion.Name)

		return &models.CommandResponse{
			OK:      true,
			Content: content,
		}
	},
}

var cmSet = models.Command{
	Command: "set",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		var idealVersion models.Version
		var idealUser models.UserPreference

		verResult := ctx.DB.Where(&models.Version{Abbreviation: params[0]}).First(&idealVersion)
		userResult := ctx.DB.Where(&models.UserPreference{UserID: ctx.UserID}).First(&idealUser)

		var response models.CommandResponse

		if verResult.Error == nil {
			if userResult.Error == nil {
				idealUser.Version = params[0]
				ctx.DB.Save(idealUser)
			} else {
				ctx.DB.Create(&models.UserPreference{
					UserID:  ctx.UserID,
					Version: params[0],
				})
			}

			response.OK = true
			response.Content = "set version"
		} else {
			response.OK = false
			response.Content = "can't find version"
		}

		return &response
	},
}
