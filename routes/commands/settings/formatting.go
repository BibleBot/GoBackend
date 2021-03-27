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
			Commands:       []models.Command{},
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

		// Make this pretty eventually
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

// Set guild prefix

// Set guild brackets (<>, [], {}, ()) <- allow more than one?

// Set user display style (default, embed, blockquote, code)

// Set user headings (true, false)

// Set user verse numbers (true, false)
