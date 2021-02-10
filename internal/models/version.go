package models

// Version is a type describing a Bible version's properties.
type Version struct {
	Name                 string
	Abbreviation         string
	Source               string
	SupportsOldTestament bool
	SupportsNewTestament bool
	SupportsDeuterocanon bool
}
