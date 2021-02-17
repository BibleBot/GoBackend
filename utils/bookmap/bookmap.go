package bookmap

import (
	_ "embed" // for go:embed
	"encoding/json"
)

//go:embed data/book_map.json
var file []byte

// GetBookmap returns the JSON object defined in names/book_map.json
func GetBookmap() map[string]map[string]string {
	var bookmap map[string]map[string]string

	json.Unmarshal(file, &bookmap)

	return bookmap
}
