package bookmap

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

// GetBookmap returns the JSON object defined in data/names/book_map.json
func GetBookmap() map[string]map[string]string {
	var bookmap map[string]map[string]string

	// If we're testing, the working directory is tests/, so paths need to be adjusted for that.
	dir := "./"
	if _, err := os.Stat(dir + "data/names/completed_names.json"); os.IsNotExist(err) {
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
