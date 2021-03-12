package converters

import (
	"encoding/json"
	"fmt"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

// InputToContext takes a byte slice and attempts to convert it to a Context model.
func InputToContext(input []byte, cfg *models.Config) (*models.Context, error) {
	var context models.Context

	err := json.Unmarshal(input, &context)
	if err != nil {
		return nil, logger.LogWithError("inputtocontext", fmt.Sprintf("failed to convert %s to Context", string(input)), err)
	}

	if context.Token != cfg.AccessKey {
		return nil, logger.LogWithError("inputtocontext", "invalid API key", nil)
	}

	context.DB = cfg.DB

	if !cfg.IsTest {
		context.DB.Where(&models.UserPreference{UserID: context.UserID}).First(&context.Prefs)
		context.DB.Where(&models.GuildPreference{GuildID: context.GuildID}).First(&context.GuildPrefs)

		if context.Prefs.Language == "" {
			if context.GuildPrefs.Language == "" {
				context.Language = cfg.Languages["english"]
			} else {
				context.Language = cfg.Languages[context.GuildPrefs.Language]
			}
		} else {
			context.Language = cfg.Languages[context.Prefs.Language]
		}

		if context.GuildPrefs.Prefix == "" {
			context.GuildPrefs.Prefix = "+"
		}
	}

	return &context, nil
}
