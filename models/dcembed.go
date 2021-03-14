package models

// DiscordEmbed is a struct representing a typical Discord RichEmbed. This is so we can output a JSONified embed that the frontend can interpret into the actual object.
type DiscordEmbed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	Colour      int     `json:"colour"`
	Footer      footer  `json:"footer"`
	Image       media   `json:"image"`
	Thumbnail   media   `json:"thumbnail"`
	Video       media   `json:"video"`
	Author      author  `json:"author"`
	Fields      []field `json:"fields"`
}

type footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type media struct {
	URL string `json:"url"`
}

type author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
