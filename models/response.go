package models

// CommandResponse is a model for command output.
type CommandResponse struct {
	OK        bool   `json:"ok"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	Pages     []page `json:"pages"`
}

type page struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// VerseResponse is a model for command output.
// TODO: Make "OK" per-verse instead of overall?
type VerseResponse struct {
	OK      bool     `json:"ok"`
	Results []*Verse `json:"verses"`
}
