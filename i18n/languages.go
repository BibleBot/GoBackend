package i18n

import (
	_ "embed" // for go:embeds
	"encoding/json"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

//go:embed english/english.json
var englishRawObjectFile []byte

// ImportLanguages imports the old JSON database into the new PGSQL one.
func ImportLanguages() []models.Language {
	var languages []models.Language

	var englishRawObject models.RawLanguage
	json.Unmarshal(englishRawObjectFile, &englishRawObject)

	englishLang := models.Language{
		Name:      "English",
		RawName:   "english",
		RawObject: englishRawObject,
	}

	languages = append(languages, englishLang)

	return languages
}
