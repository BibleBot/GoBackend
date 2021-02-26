package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/textpurification"
)

// APIBibleProvider is a Provider for API.Bible-based versions.
type APIBibleProvider struct {
	Key string
}

var (
	abOnce       sync.Once
	abInstance   *APIBibleProvider
	versionTable = map[string]string{
		"KJVA": "de4e12af7f28f599-01",
		"FBV":  "65eec8e0b60e656b-01",
	}
)

// NewAPIBibleProvider creates a APIBibleProvider if one does not already exist, otherwise returns existing instance.
func NewAPIBibleProvider(key string) *APIBibleProvider {
	abOnce.Do(func() {
		abInstance = &APIBibleProvider{Key: key}
	})

	return abInstance
}

// GetVerse fetches a Verse based upon a Reference, for API.Bible versions.
func (abp *APIBibleProvider) GetVerse(ref *models.Reference, titles bool, verseNumbers bool) (*models.Verse, error) {
	URL, err := url.Parse(fmt.Sprintf("https://api.scripture.api.bible/v1/bibles/%s/search", versionTable[ref.Version.Abbreviation]))
	if err != nil {
		return nil, logger.LogWithError("apibible", "unable to create URL", err)
	}

	query := URL.Query()
	query.Set("query", ref.ToString())
	query.Set("limit", "1")
	URL.RawQuery = query.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, logger.LogWithError("apibible", "unable to create request", err)
	}
	req.Header.Add("api-key", abp.Key)

	resp, err := client.Do(req)
	if err != nil {
		return nil, logger.LogWithError("apibible", "unable to fetch verse", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, logger.LogWithError("apibible", "unable to fetch verse - did not return status 200", nil)
	}

	abSearchResponse := new(models.ABSearchResponse)
	json.NewDecoder(resp.Body).Decode(abSearchResponse)

	data := abSearchResponse.Data.Passages

	if len(data) == 0 {
		return nil, logger.LogWithError("apibible", "no passages found", nil)
	}

	if data[0].BibleID != versionTable[ref.Version.Abbreviation] {
		return nil, logger.LogWithError("apibible", fmt.Sprintf("%s is no longer able to be used", ref.Version.Abbreviation), nil)
	}

	if data[0].Content == "" {
		return nil, logger.LogWithError("apibible", "passage contains no content", nil)
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(data[0].Content))
	if err != nil {
		return nil, logger.LogWithError("apibible", "unable to parse document", err)
	}

	document.Find(".v").Each(func(idx int, element *goquery.Selection) {
		if verseNumbers {
			element.SetText(fmt.Sprintf("<**%s**> ", element.Text()))
		} else {
			element.Remove()
		}
	})

	document.Find(".nd").Each(func(idx int, element *goquery.Selection) {
		if element.Text() == "LORD" {
			element.SetText("Lᴏʀᴅ")
		} else {
			var newText []rune
			for _, char := range element.Text() {
				newChar := char
				idx := slices.IndexRune(alphabet, char)

				if idx != -1 {
					for sIdx, sChar := range smallcaps {
						if sIdx == idx {
							newChar = sChar
						}
					}
				}

				newText = append(newText, newChar)
			}
			element.SetText(string(newText))
		}
	})

	title := ""
	if titles {
		title = strings.Join(document.Find("h3").Map(func(idx int, element *goquery.Selection) string {
			return strings.TrimSpace(element.Text())
		}), " / ")
	}

	text := strings.Join(document.Find("p").Map(func(idx int, element *goquery.Selection) string {
		return strings.TrimSpace(element.Text())
	}), "\n")

	return &models.Verse{
		Reference: ref,
		Title:     title,
		Text:      textpurification.PurifyVerseText(text),
	}, nil
}

// Search gathers search results based on a query.
func (abp APIBibleProvider) Search(query string, version *models.Version) (*map[string]string, error) {
	return nil, nil
}
