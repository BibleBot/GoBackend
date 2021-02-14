package models

// Verse is a type containing a verse's content including any applicable references and versions.
type Verse struct {
	Reference *Reference
	Title     string
	Text      string
}
