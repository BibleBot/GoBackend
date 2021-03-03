package models

import "gorm.io/gorm"

// Context is a model outlining commonly-passed parameters.
type Context struct {
	DB         gorm.DB
	Prefs      UserPreference
	GuildPrefs GuildPreference
	Language   Language

	Body        string `json:"body"`
	TempVersion string `json:"ver"`
}
