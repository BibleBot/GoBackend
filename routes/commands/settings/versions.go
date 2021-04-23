package settings

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/embedify"
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
			DefaultCommand: verDefault,
			Commands:       []models.Command{verSet, verSetServer, verList},
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
		}

		return false
	}).([]models.Command)

	var cm models.Command
	if len(cmMatches) == 0 {
		cm = cr.DefaultCommand
	} else {
		cm = cmMatches[0]
	}

	if len(params) == 0 {
		return cm.Process([]string{}, ctx)
	}

	return cm.Process(params[1:], ctx)
}

var verDefault = models.Command{
	Command: "version",
	Process: func(_ []string, ctx *models.Context) *models.CommandResponse {
		lng := ctx.Language

		var userVersion models.Version
		var guildVersion models.Version

		if ctx.Prefs.Version != "" {
			ctx.DB.Where(&models.Version{Abbreviation: ctx.Prefs.Version}).First(&userVersion)
		}

		if ctx.GuildPrefs.Version != "" {
			ctx.DB.Where(&models.Version{Abbreviation: ctx.GuildPrefs.Version}).First(&guildVersion)
		}

		tVersionUsed := strings.ReplaceAll(lng.GetRawString("VersionUsed"), "<version>", "**"+userVersion.Name+"**")
		tGuildVersionUsed := strings.ReplaceAll(lng.GetRawString("ServerVersionUsed"), "<version>", "**"+guildVersion.Name+"**")

		tUsage := "**%s** - %s"
		tSetVersionUsage := fmt.Sprintf(tUsage, lng.GetString("set"), lng.GetString("SetVersionUsage"))
		tSetGuildVersionUsage := fmt.Sprintf(tUsage, lng.GetString("setserver"), lng.GetString("SetServerVersionUsage"))
		tInfoUsage := fmt.Sprintf(tUsage, lng.GetString("info"), lng.GetString("InfoUsage"))
		tListUsage := fmt.Sprintf(tUsage, lng.GetString("list"), lng.GetString("ListVersionUsage"))

		content := fmt.Sprintf("%s\n%s\n\n__%s__\n%s\n%s\n%s\n%s",
			tVersionUsed, tGuildVersionUsed, lng.GetString("Subcommands"),
			tSetVersionUsage, tSetGuildVersionUsage, tInfoUsage, tListUsage)

		return &models.CommandResponse{
			OK:      true,
			Content: embedify.Embedify("", "+version", content, false, ""),
		}
	},
}

var verSet = models.Command{
	Command: "set",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		lng := ctx.Language

		var idealVersion models.Version
		var idealUser models.UserPreference

		verResult := ctx.DB.Where(&models.Version{Abbreviation: params[0]}).First(&idealVersion)
		userResult := ctx.DB.Where(&models.UserPreference{UserID: ctx.UserID}).First(&idealUser)

		var response models.CommandResponse

		if verResult.Error == nil && params[0] == idealVersion.Abbreviation {
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
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString("<+><version> <set>"), lng.GetString("SetVersionSuccess"), false, "")
		} else {
			response.OK = false
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString("<+><version> <set>"), lng.GetString("SetVersionFail"), false, "")
		}

		return &response
	},
}

var verSetServer = models.Command{
	Command: "setserver",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		lng := ctx.Language

		var idealVersion models.Version
		var idealGuild models.GuildPreference

		verResult := ctx.DB.Where(&models.Version{Abbreviation: params[0]}).First(&idealVersion)
		guildResult := ctx.DB.Where(&models.GuildPreference{GuildID: ctx.GuildID}).First(&idealGuild)

		var response models.CommandResponse

		if verResult.Error == nil && params[0] == idealVersion.Abbreviation {
			if guildResult.Error == nil {
				idealGuild.Version = params[0]
				ctx.DB.Save(idealGuild)
			} else {
				ctx.DB.Create(&models.GuildPreference{
					GuildID: ctx.GuildID,
					Version: params[0],
				})
			}

			response.OK = true
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString("<+><version> <setserver>"), lng.GetString("SetServerVersionSuccess"), false, "")
		} else {
			response.OK = false
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString("<+><version> <setserver>"), lng.GetString("SetServerVersionFail"), false, "")
		}

		return &response
	},
}

var verList = models.Command{
	Command: "list",
	Process: func(_ []string, ctx *models.Context) *models.CommandResponse {
		var versions []models.Version
		ctx.DB.Find(&versions)

		// Sort versions alphabetically.
		sort.SliceStable(versions, func(i, j int) bool {
			return versions[i].Name < versions[j].Name
		})

		var pages []*models.DiscordEmbed
		var maxResultsPerPage = 25
		var versionsUsed []models.Version
		totalPages := int(math.Ceil(float64(len(versions)) / float64(maxResultsPerPage)))

		if totalPages == 0 {
			totalPages = 1
		}

		for i := 0; i < totalPages; i++ {
			pageCounter := ctx.Language.GetString("PageOf")
			pageCounter = strings.ReplaceAll(pageCounter, "<num>", strconv.Itoa(i+1))
			pageCounter = strings.ReplaceAll(pageCounter, "<total>", strconv.Itoa(totalPages))

			title := fmt.Sprintf("%s - %s", ctx.Language.TranslatePlaceholdersInString("<+><version> <list>"), pageCounter)
			embed := embedify.Embedify("", title, "", false, "")

			count := 0
			versionList := ""

			for _, version := range versions {
				if count < maxResultsPerPage {
					if !slices.VersionInSlice(version, versionsUsed) {
						versionList += fmt.Sprintf("%s\n", version.Name)

						versionsUsed = append(versionsUsed, version)
						versions = versions[1:]
						count++
					}
				}
			}

			embed.Description = versionList
			pages = append(pages, embed)
		}

		return &models.CommandResponse{
			OK:       true,
			Language: &ctx.Language,
			Pages:    pages,
		}
	},
}
