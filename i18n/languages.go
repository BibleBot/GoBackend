package i18n

import (
	_ "embed" // for go:embeds
	"encoding/json"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

//go:embed english/english.json
var englishRawObjectFile []byte

// ImportLanguages simply imports any defined languages and returns them in a map.
func ImportLanguages() map[string]models.Language {
	languages := make(map[string]models.Language)

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
