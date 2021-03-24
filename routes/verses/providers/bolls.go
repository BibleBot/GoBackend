package providers

import (
	"bytes"
	_ "embed" // for go:embed
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/textpurification"
)

// BollsProvider is a Provider for Bolls Bible-based versions.
type BollsProvider struct {
	Key        string
	NamesToIDs map[string]int
}

type bollsBody struct {
	Translations []string `json:"translations"`
	Verses       []int    `json:"verses"`
	Book         int      `json:"book"`
	Chapter      int      `json:"chapter"`
}

var (
	boOnce     sync.Once
	boInstance *BollsProvider

	//go:embed data/bolls.json
	nameToIDFile []byte
)

// NewbollsProvider creates a BollsProvider if one does not already exist, otherwise returns existing instance.
func NewBollsProvider() *BollsProvider {
	boOnce.Do(func() {
		namesToIDs := make(map[string]int)
		json.Unmarshal(nameToIDFile, &namesToIDs)

		boInstance = &BollsProvider{NamesToIDs: namesToIDs}
	})

	return boInstance
}

// GetVerse fetches a Verse based upon a Reference, for Bolls Bible versions.
func (bop *BollsProvider) GetVerse(ref *models.Reference, titles bool, verseNumbers bool) (*models.Verse, error) {
	var verses []int
	for i := ref.StartingVerse; i < ref.EndingVerse; i++ {
		verses = append(verses, i)
	}

	bodyBytes, _ := json.Marshal(bollsBody{
		Translations: []string{"YLT"},
		Verses:       verses,
		Book:         bop.NamesToIDs[ref.Book],
		Chapter:      ref.StartingChapter,
	})
	body := bytes.NewReader(bodyBytes)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://bolls.life/get-paralel-verses/", body)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, logger.LogWithError("bolls", "unable to fetch verse", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, logger.LogWithError("bolls", "unable to fetch verse - did not return status 200", nil)
	}

	abSearchResponse := new(models.ABSearchResponse)
	json.NewDecoder(resp.Body).Decode(abSearchResponse)

	data := abSearchResponse.Data.Passages

	if len(data) == 0 {
		return nil, logger.LogWithError("bolls", "no passages found", nil)
	}

	if data[0].BibleID != versionTable[ref.Version.Abbreviation] {
		return nil, logger.LogWithError("bolls", fmt.Sprintf("%s is no longer able to be used", ref.Version.Abbreviation), nil)
	}

	if data[0].Content == "" {
		return nil, logger.LogWithError("bolls", "passage contains no content", nil)
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(data[0].Content))
	if err != nil {
		return nil, logger.LogWithError("bolls", "unable to parse document", err)
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
func (bop BollsProvider) Search(query string, version *models.Version) (*map[string]string, error) {
	return nil, nil
}
