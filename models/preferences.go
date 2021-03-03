package models

import "gorm.io/gorm"

// UserPreference is a model for user-based preferences.
type UserPreference struct {
	gorm.Model
	UserID       string `json:"id"`
	Input        string `json:"input"`
	Language     string `json:"language"`
	Version      string `json:"version"`
	Titles       bool   `json:"fmtTitles"`
	VerseNumbers bool   `json:"fmtVerseNumbers"`
	DisplayMode  string `json:"fmtDisplayMode"`
}

// GuildPreference is a model for guild/DM-based preferences.
type GuildPreference struct {
	gorm.Model
	GuildID          string `json:"id"`
	Language         string `json:"language"`
	Version          string `json:"version"`
	Prefix           string `json:"fmtPrefix"`
	IgnoringBrackets string `json:"fmtIgnoringBrackets"`
	DVChannel        string `json:"dvChannel"`
	DVTime           string `json:"dvTime"`
	DVTimeZone       string `json:"dvTimeZone"`
	IsDM             bool   `json:"isDM"`
}
