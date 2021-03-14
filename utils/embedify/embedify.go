package embedify

import (
	"fmt"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

const (
	NORMAL_COLOR = 6709986
	ERROR_COLOR  = 16723502
)

func Embedify(author string, title string, description string, isError bool, copyright string) (*models.DiscordEmbed, error) {
	var embed models.DiscordEmbed

	embed.Colour = NORMAL_COLOR
	if isError {
		embed.Colour = ERROR_COLOR
	}

	embed.Footer.Text = "BibleBot v9.1.0"                    // TODO: Make this not hard-coded.
	embed.Footer.IconURL = "https://i.imgur.com/hr4RXpy.png" // Make this also not hard-coded.

	if title != "" {
		embed.Title = title
	}

	if author != "" {
		embed.Author.Name = author
	}

	if description != "" {
		if len(description) < 2049 {
			embed.Description = description
		} else {
			return nil, logger.LogWithError("embedify", "description is too long", nil)
		}
	}

	if copyright != "" {
		embed.Footer.Text = fmt.Sprintf("%s // %s", copyright, embed.Footer.Text)
	}

	return &embed, nil
}
