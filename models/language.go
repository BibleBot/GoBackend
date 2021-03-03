package models

// Language is a type describing an interface language.
type Language struct {
	Name      string
	RawName   string
	RawObject rawLanguage
}

type rawLanguage struct {
	// TODO
}
