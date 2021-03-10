package i18n

import (
	_ "embed" // for go:embeds
	"encoding/json"

	"gorm.io/gorm"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

// ImportLanguages imports the old JSON database into the new PGSQL one.
func ImportLanguages(db *gorm.DB) {

	//go:embed english/english.json
	var englishRawObjectFile []byte
	var englishRawObject models.RawLanguage
	json.Unmarshal(englishRawObjectFile, &englishRawObject)

	englishLang := &models.Language{
		Name:      "English",
		RawName:   "english",
		RawObject: englishRawObject,
	}

	if err := db.Where(englishLang).First(&models.Language{}).Error; err != nil {
		db.Create(englishLang)
	}
}
