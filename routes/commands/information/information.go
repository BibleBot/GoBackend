package information

import (
	"fmt"
	"strings"
	"sync"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/embedify"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

// InformationCommandRouter is a basic struct with functions to handle help-related commands.
type InformationCommandRouter struct {
	DefaultCommand models.Command
	Commands       []models.Command
}

var (
	// infoInstance is the singleton router used to process its respective commands
	infoInstance *InformationCommandRouter
	infoOnce     sync.Once
)

// NewInformationCommandRouter creates a HelpCommandRouter if one does not already exist
func NewInformationCommandRouter() *InformationCommandRouter {
	infoOnce.Do(func() {
		infoInstance = &InformationCommandRouter{
			DefaultCommand: infoBibleBot,
			Commands:       []models.Command{},
		}
	})

	return infoInstance
}

// Process checks which command process to run given the inputed command & parameters
func (cr *InformationCommandRouter) Process(params []string, ctx *models.Context) *models.CommandResponse {
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

var infoBibleBot = models.Command{
	Command: "biblebot",
	Process: func(params []string, ctx *models.Context) *models.CommandResponse {
		if len(params) > 0 {
			// TODO: implement rescue option
			return &models.CommandResponse{}
		} else {
			lng := ctx.Language
			title := strings.Replace(lng.GetRawString("BibleBot"), "<version>", ctx.Version, 1)
			desc := lng.GetString(ctx, "Credit")

			commandListName := lng.GetString(ctx, "CommandListName")
			commandList := lng.GetString(ctx, "CommandList")

			linksName := lng.GetString(ctx, "Links")
			links := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n\n**%s**", lng.GetRawString("Website"), lng.GetRawString("Copyrights"),
				lng.GetRawString("Code"), lng.GetRawString("Server"),
				lng.GetRawString("Terms"), lng.GetRawString("Usage"))

			replacementMap := map[string]string{
				"<website>":    "https://biblebot.xyz",
				"<copyrights>": "https://biblebot.xyz/copyrights",
				"<repository>": "https://internal.kerygma.digital",
				"<invite>":     "https://discord.gg/H7ZyHqE",
				"<terms>":      "https://biblebot.xyz/terms",
			}

			for k, v := range replacementMap {
				links = strings.ReplaceAll(links, k, v)
			}

			embed := embedify.Embedify("", title, desc, false, "")
			embed.Fields = append(embed.Fields, models.EmbedField{
				Name:   commandListName,
				Value:  commandList,
				Inline: false,
			})
			embed.Fields = append(embed.Fields, models.EmbedField{
				Name:   "\u200B",
				Value:  "—————————————",
				Inline: false,
			})
			embed.Fields = append(embed.Fields, models.EmbedField{
				Name:   linksName,
				Value:  links,
				Inline: false,
			})

			return &models.CommandResponse{
				OK:      true,
				Content: embed,
			}
		}
	},
}
