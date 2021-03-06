package dbimports

import (
	_ "embed" // for go:embed	"encoding/json"
	"encoding/json"

	"gorm.io/gorm"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
)

var (
	//go:embed data/versiondb.json
	oldVersions []byte
)

type oldVersion struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbv"`
	HasOT        bool   `json:"hasOT"`
	HasNT        bool   `json:"hasNT"`
	HasDEU       bool   `json:"hasDEU"`
}

// ImportVersions imports the old JSON database into the new PGSQL one.
func ImportVersions(db *gorm.DB) {
	var oldVersionEntries map[string]oldVersion
	json.Unmarshal(oldVersions, &oldVersionEntries)

	for _, val := range oldVersionEntries {
		source := "bg"

		if slices.StringInSlice(val.Abbreviation, []string{"BSB", "NHEB", "WBT"}) {
			continue
		} else if slices.StringInSlice(val.Abbreviation, []string{"KJVA", "FBV"}) {
			source = "ab"
		} else if slices.StringInSlice(val.Abbreviation, []string{"ELXX", "LXX"}) {
			continue
		}

		newVersion := &models.Version{
			Name:                 val.Name,
			Abbreviation:         val.Abbreviation,
			Source:               source,
			SupportsOldTestament: val.HasOT,
			SupportsNewTestament: val.HasNT,
			SupportsDeuterocanon: val.HasDEU,
		}

		if err := db.Where(newVersion).First(&models.Version{}).Error; err != nil {
			db.Create(newVersion)
		}
	}
}
