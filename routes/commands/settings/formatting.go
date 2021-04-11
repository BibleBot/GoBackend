package settings

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/embedify"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// FormatCommandRouter is a basic struct with functions to handle format-related commands.
type FormatCommandRouter struct {
	DefaultCommand models.Command
	Commands       []models.Command
}

var (
	// formattingInstance is the singleton router used to process its respective commands
	formattingInstance *FormatCommandRouter
	formattingOnce     sync.Once
)

// NewFormattingCommandRouter creates a FormatCommandRouter if one does not already exist
func NewFormattingCommandRouter() *FormatCommandRouter {
	formattingOnce.Do(func() {
		formattingInstance = &FormatCommandRouter{
			DefaultCommand: fmtDefault,
			Commands:       []models.Command{fmtPrefix, fmtBrackets, fmtStyle, fmtHeadings, fmtVerseNumbers},
		}
	})

	return formattingInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *FormatCommandRouter) Process(params []string, ctx *models.Context) *models.CommandResponse {
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

// TODO: Improve/Check
// Get user preferences (display style, headings, verse numbers), guild preferences (prefix, brackets), and a list of formatting commands
var fmtDefault = models.Command{
	Command: "formatting",
	Process: func(_ []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		lng := ctx.Language

		guildPrefix := ctx.GuildPrefs.Prefix
		userDisplayStyle := ctx.Prefs.DisplayMode
		userHeadings := ctx.Prefs.Titles
		userVerseNumbers := ctx.Prefs.VerseNumbers
		guildBrackets := ctx.GuildPrefs.IgnoringBrackets

		// Check if not nil?

		var response models.CommandResponse

		// Make this pretty eventually (simplify from full sentences in i18n to simple Key: Value?)
		content := fmt.Sprintf("Current Preferences:\n%s\n%s\n%s\n%s\n%s\n\nFormatting Commands:\n",
			strings.Replace(lng.GetString(ctx, "GuildPrefixUsed"), "<prefix>", guildPrefix, 1),
			strings.Replace(lng.GetString(ctx, "Formatting"), "<value>", userDisplayStyle, 1),
			strings.Replace(lng.GetString(ctx, "Headings"), "<status>", strconv.FormatBool(userHeadings), 1),
			strings.Replace(lng.GetString(ctx, "VerseNumbers"), "<status>", strconv.FormatBool(userVerseNumbers), 1),
			strings.Replace(lng.GetString(ctx, "GuildBracketsUsed"), "<brackets>", guildBrackets, 1),
		)

		response.OK = true
		response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString(ctx, "<+><formatting>"), content, false, "")

		return &response
	},
}

// TODO: Improve/Check
// Set guild prefix
var fmtPrefix = models.Command{
	Command: "prefix",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		// if not owner/manage permissions return error

		fmt.Println(ctx.Prefs)
		lng := ctx.Language

		var idealGuild models.GuildPreference

		guildResult := ctx.DB.Where(&models.GuildPreference{GuildID: ctx.GuildID}).First(&idealGuild)

		var response models.CommandResponse

		if len(params[0]) == 1 {
			if guildResult.Error == nil {
				idealGuild.Prefix = params[0]
				ctx.DB.Save(idealGuild)
			} else {
				ctx.DB.Create(&models.GuildPreference{
					GuildID: ctx.GuildID,
					Prefix:  params[0],
				})
			}

			response.OK = true
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString(ctx, "<+><formatting> <setprefix>"), lng.GetString(ctx, "PrefixSuccess"), false, "")
		} else {
			response.OK = false
			response.Content = embedify.Embedify("", lng.TranslatePlaceholdersInString(ctx, "<+><formatting> <setprefix>"), lng.GetString(ctx, "PrefixOneChar"), false, "")
		}

		return &response
	},
}

// TODO: Implement
// Set guild brackets (<>, [], {}, ()) <- allow more than one?
var fmtBrackets = models.Command{
	Command: "brackets",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		// lng := ctx.Language

		return nil
	},
}

// TODO: Implement
// Set user display style (default, embed, blockquote, code)
var fmtStyle = models.Command{
	Command: "style",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		// lng := ctx.Language

		return nil
	},
}

// TODO: Implement
// Set user headings (true, false)
var fmtHeadings = models.Command{
	Command: "headings",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		// lng := ctx.Language

		return nil
	},
}

// TODO: Implement
// Set user verse numbers (true, false)
var fmtVerseNumbers = models.Command{
	Command: "numbers",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		fmt.Println(ctx.Prefs)
		// lng := ctx.Language

		return nil
	},
}
