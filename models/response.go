package models

// CommandResponse is a model for command output.
type CommandResponse struct {
	OK       bool            `json:"ok"`
	Language *Language       `json:"language"`
	Content  *DiscordEmbed   `json:"content"`
	Pages    []*DiscordEmbed `json:"pages"`
}

// VerseResponse is a model for command output.
// TODO: Make "OK" per-verse instead of overall?
type VerseResponse struct {
	OK      bool     `json:"ok"`
	Results []*Verse `json:"verses"`
}
