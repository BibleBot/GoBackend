package models

// Provider is a basic interface defining functions expected in Bible providers.
type Provider interface {
	GetVerse(*Reference, bool, bool) (*Verse, error)
	Search(string, *Version) (*map[string]string, error)
}
