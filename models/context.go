package models

import "gorm.io/gorm"

// Context is a model outlining commonly-passed parameters.
type Context struct {
	DB         gorm.DB
	Prefs      UserPreference
	GuildPrefs GuildPreference
	Language   Language

	Token string `json:"token"`
	Body  string `json:"body"`
	IsDM  bool   `json:"isDM"`

	UserID    string `json:"userID"`
	ChannelID string `json:"channelID"`
	GuildID   string `json:"guildID"`
	Version   string `json:"version"`
}
