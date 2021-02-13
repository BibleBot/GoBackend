package namefetcher

// ABBibleResponse is a struct corresponding to the response from /v1/bibles.
type ABBibleResponse struct {
	Data []abBibleData `json:"data"`
}

// ABBookResponse is a struct corresponding to the response from /v1/bibles/{bibleId}/books.
type ABBookResponse struct {
	Data []abBookData `json:"data"`
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
