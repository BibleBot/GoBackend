package models

// ABBibleResponse is a struct corresponding to the response from /v1/bibles.
type ABBibleResponse struct {
	Data []abBibleData `json:"data"`
}

// ABBookResponse is a struct corresponding to the response from /v1/bibles/{bibleId}/books.
type ABBookResponse struct {
	Data []abBookData `json:"data"`
}

// ABSearchResponse is a struct corresponding to the response from /v1/bibles/{bibleId}/search.
type ABSearchResponse struct {
	Query string       `json:"query"`
	Data  abSearchData `json:"data"`
	Meta  abMetadata   `json:"meta"`
}

type abBookData struct {
	ID           string      `json:"id"`
	BibleID      string      `json:"bibleId"`
	Abbreviation string      `json:"abbreviation"`
	Name         string      `json:"name"`
	NameLong     string      `json:"nameLong"`
	Chapters     []abChapter `json:"chapters"`
}

type abChapter struct {
	ID        string `json:"id"`
	BibleID   string `json:"bibleId"`
	Number    string `json:"number"` // no, really, it's a string
	BookID    string `json:"bookId"`
	Reference string `json:"reference"`
}

type abBibleData struct {
	ID                string         `json:"id"`
	DBLID             string         `json:"dblId"`
	Abbreviation      string         `json:"abbreviation"`
	AbbreviationLocal string         `json:"abbreviationLocal"`
	Language          abLanguage     `json:"language"`
	Countries         []abCountry    `json:"countries"`
	Name              string         `json:"name"`
	NameLocal         string         `json:"nameLocal"`
	Description       string         `json:"description"`
	DescriptionLocal  string         `json:"descriptionLocal"`
	RelatedDBL        string         `json:"relatedDbl"`
	Type              string         `json:"type"`
	UpdatedAt         string         `json:"updatedAt"`
	AudioBibles       []abAudioBible `json:"audioBibles"`
}

type abLanguage struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	NameLocal       string `json:"nameLocal"`
	Script          string `json:"script"`
	ScriptDirection string `json:"scriptDirection"`
}

type abCountry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	NameLocal string `json:"nameLocal"`
}

type abAudioBible struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	NameLocal        string `json:"nameLocal"`
	Description      string `json:"description"`
	DescriptionLocal string `json:"descriptionLocal"`
}

type abSearchData struct {
	Query      string      `json:"query"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Total      int         `json:"total"`
	VerseCount int         `json:"verseCount"`
	Verses     []abVerse   `json:"verses"`
	Passages   []abPassage `json:"passages"`
}

type abVerse struct {
	ID        string `json:"id"`
	OrgID     string `json:"orgId"`
	BibleID   string `json:"bibleId"`
	BookID    string `json:"bookId"`
	ChapterID string `json:"chapterId"`
	Text      string `json:"text"`
	Reference string `json:"reference"`
}

type abPassage struct {
	ID         string `json:"id"`
	BibleID    string `json:"bibleId"`
	OrgID      string `json:"orgId"`
	Content    string `json:"content"`
	Reference  string `json:"reference"`
	VerseCount int    `json:"verseCount"`
	Copyright  string `json:"copyright"`
}

type abMetadata struct {
	FUMS          string `json:"fums"`
	FUMSID        string `json:"fumsId"`
	FUMSJSInclude string `json:"fumsJsInclude"`
	FUMSJS        string `json:"fumsJs"`
	FUMSNoScript  string `json:"fumsNoScript"`
}
