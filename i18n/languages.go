package i18n

import (
	_ "embed" // for go:embeds
	"encoding/json"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

//go:embed english/english.json
var englishRawObjectFile []byte

// ImportLanguages imports the old JSON database into the new PGSQL one.
func ImportLanguages() map[string]models.Language {
	var languages map[string]models.Language

	var englishRawObject models.RawLanguage
	json.Unmarshal(englishRawObjectFile, &englishRawObject)

	englishLang := models.Language{
		Name:      "English",
		RawName:   "english",
		RawObject: englishRawObject,
	}

	languages["english"] = englishLang

	return languages
}
