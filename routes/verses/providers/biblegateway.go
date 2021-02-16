package providers

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/textpurification"
)

// BibleGatewayProvider is a Provider for BibleGateway-based versions.
type BibleGatewayProvider struct{}

var (
	once     sync.Once
	instance *BibleGatewayProvider
)

// NewBibleGatewayProvider creates a BibleGatewayProvider if one does not already exist, otherwise returns existing instance.
func NewBibleGatewayProvider() *BibleGatewayProvider {
	once.Do(func() {
		instance = &BibleGatewayProvider{}
	})

	return instance
}

// GetVerse fetches a Verse based upon a Reference, for BibleGateway versions.
func (bgp *BibleGatewayProvider) GetVerse(ref *models.Reference, titles bool, verseNumbers bool) (*models.Verse, error) {
	URL := fmt.Sprintf("https://www.biblegateway.com/passage/?search=%s&version=%s&interface=print", ref.ToString(), ref.Version.Abbreviation)

	resp, err := http.Get(URL)
	if err != nil {
		return nil, logger.LogWithError("biblegateway", "unable to fetch verse", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, logger.LogWithError("biblegateway", "unable to fetch verse - did not return status 200", nil)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, logger.LogWithError("biblegateway", "unable to parse document", err)
	}
	container := document.Find(".passage-col").First()

	container.Find(".chapternum").Each(func(idx int, element *goquery.Selection) {
		if verseNumbers {
			element.SetText("<**1**> ")
		} else {
			element.Remove()
		}
	})

	container.Find(".versenum").Each(func(idx int, element *goquery.Selection) {
		if verseNumbers {
			element.SetText(fmt.Sprintf("<**%s**> ", element.Text()[:len(element.Text())-2]))
		} else {
			element.Remove()
		}
	})

	container.Find("br").Each(func(idx int, element *goquery.Selection) {
		element.SetHtml("\n")
	})

	container.Find(".crossreference").Each(func(idx int, element *goquery.Selection) {
		element.Remove()
	})

	container.Find(".footnote").Each(func(idx int, element *goquery.Selection) {
		element.Remove()
	})

	title := ""
	if titles {
		title = strings.Join(container.Find("h3").Map(func(idx int, element *goquery.Selection) string {
			return strings.TrimSpace(element.Text())
		}), " / ")
	}

	text := strings.Join(container.Find("p").Map(func(idx int, element *goquery.Selection) string {
		return strings.TrimSpace(element.Text())
	}), "\n")

	return &models.Verse{
		Reference: ref,
		Title:     title,
		Text:      textpurification.PurifyVerseText(text),
	}, nil
}

// Search gathers search results based on a query.
func (bgp BibleGatewayProvider) Search(query string, version *models.Version) (*map[string]string, error) {
	return nil, nil
}
