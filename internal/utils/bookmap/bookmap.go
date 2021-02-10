package bookmap

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BibleBot/backend/internal/utils/logger"
)

// GetBookmap returns the JSON object defined in data/names/book_map.json
func GetBookmap(isTest bool) map[string]map[string]string {
	var bookmap map[string]map[string]string

	// If we're testing, the working directory is tests/, so paths need to be adjusted for that.
	dir := "./"
	if isTest {
		dir = "./../"
	}

	file, err := ioutil.ReadFile(dir + "data/names/book_map.json")
	if err != nil {
		logger.LogWithError("namefetcher", "failed to open book_map.json, run backend normally before testing again", err)
		os.Exit(3)
	}
	json.Unmarshal(file, &bookmap)

	return bookmap
}
