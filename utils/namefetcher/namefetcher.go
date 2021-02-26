package namefetcher

import (
	_ "embed" // for go:embed

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/slices"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	//go:embed data/apibible_names.json
	apiBibleFile []byte

	//go:embed data/abbreviations.json
	abbreviationsFile []byte

	//go:embed data/default_names.json
	defaultNamesFile []byte

	//go:embed data/nuisances.json
	nuisancesFile []byte

	apiBibleNames = make(map[string]string)
	abbreviations = make(map[string][]string)
	bookNames     = make(map[string][]string)
	defaultNames  []string
	nuisances     []string
)

// GetBookNames returns map[string][]string of saved book names.
func GetBookNames() map[string][]string {
	if len(bookNames) == 0 {
		// Get mapping of API.Bible books to BibleGateway, which we use as a standard.
		file, err := ioutil.ReadFile("./utils/namefetcher/data/completed_names.json")
		if err != nil {
			logger.LogWithError("namefetcher", "failed to open completed_names.json, run backend normally before testing again", err)
			os.Exit(3)
		}

		json.Unmarshal(file, &bookNames)
	}

	return bookNames
}

// GetDefaultBookNames returns []string of default book names.
func GetDefaultBookNames() []string {
	if len(defaultNames) == 0 {
		json.Unmarshal(defaultNamesFile, &defaultNames)
	}

	return defaultNames
}

// FetchBookNames goes through all of BibleGateway and API.Bible, scraping book names from each translation.
func FetchBookNames(apiBibleKey string, isDryRun bool, isTest bool) error {
	// Create a spinner, including our usual log prefixes.
	hiCyan := color.New(color.FgHiCyan).SprintFunc()
	hiMagenta := color.New(color.FgHiMagenta).SprintFunc()
	sp := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sp.Prefix = hiCyan("[info] ") + hiMagenta("<namefetcher> ")

	// We do not want to run on dry runs or testing.
	if isDryRun && !isTest {
		sp.FinalMSG = hiCyan("[info] ") + hiMagenta("<namefetcher> ") + "✔️  Name fetching set to dry, skipping.\n"

		sp.Start()
		sp.Stop()

		return nil
	} else if isTest {
		sp.FinalMSG = hiCyan("[info] ") + hiMagenta("<namefetcher> ") + "❌  Name fetching is disabled for tests, skipping.\n"

		return nil
	}

	// Get mapping of API.Bible books to BibleGateway, which we use as a standard.
	json.Unmarshal(apiBibleFile, &apiBibleNames)

	// Get standard English abbreviations.
	json.Unmarshal(abbreviationsFile, &abbreviations)

	// Get standard book IDs.
	json.Unmarshal(defaultNamesFile, &defaultNames)

	// Pre-flight checks have cleared. Houston, we have liftoff.
	bgVersions, err := getBibleGatewayVersions(sp)
	if err != nil {
		return err
	}

	bgNames, err := getBibleGatewayNames(bgVersions, sp)
	if err != nil {
		return err
	}

	abVersions, err := getAPIBibleVersions(apiBibleKey, sp)
	if err != nil {
		return err
	}

	abNames, err := getAPIBibleNames(abVersions, apiBibleKey, sp)
	if err != nil {
		return err
	}

	sp.Suffix = "  Writing to file..."

	_, err = os.Stat("./utils/namefetcher/data/completed_names.json")
	if !os.IsNotExist(err) {
		err = os.Remove("./utils/namefetcher/data/completed_names.json")

		if err != nil {
			sp.Stop()
			return logger.LogWithError("namefetcher", "failed to remove completed_names.json, invalid permissions?", err)
		}
	}

	completedNames := mergeThreeMaps(bgNames, abNames, abbreviations)

	resultFile, err := os.OpenFile("./utils/namefetcher/data/completed_names.json", os.O_CREATE, os.ModePerm)
	if err != nil {
		sp.Stop()
		return logger.LogWithError("namefetcher", "failed to open completed_names.json", err)
	}
	defer resultFile.Close()

	jsonEncoder := json.NewEncoder(resultFile)
	err = jsonEncoder.Encode(completedNames)
	if err != nil {
		sp.Stop()
		return logger.LogWithError("namefetcher", "failed to write completed_names.json", err)
	}

	sp.FinalMSG = hiCyan("[info] ") + hiMagenta("<namefetcher> ") + "✔️  Name fetcher finished, wrote file successfully.\n"
	sp.Stop()

	return nil
}

func getBibleGatewayVersions(sp *spinner.Spinner) (map[string]string, error) {
	sp.Suffix = "  Fetching BibleGateway versions..."
	sp.Start()

	versions := make(map[string]string)

	resp, err := http.Get("https://www.biblegateway.com/versions/")
	if err != nil {
		return nil, logger.LogWithError("namefetcher", "couldn't reach biblegateway version list", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("biblegateway version list did not respond 200, got %d", resp.StatusCode)
		return nil, logger.LogWithError("namefetcher", msg, err)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, logger.LogWithError("namefetcher", "couldn't read biblegateway version list", err)
	}

	document.Find(".translation-name").Each(func(idx int, element *goquery.Selection) {
		target := element.Find("a")

		text := target.Text()
		link, exists := target.Attr("href")

		if exists {
			versions[text] = fmt.Sprintf("https://www.biblegateway.com%s", link)
		}
	})

	return versions, nil
}

func getBibleGatewayNames(versions map[string]string, sp *spinner.Spinner) (map[string][]string, error) {
	names := make(map[string][]string)

	for versionName, versionLink := range versions {
		sp.Suffix = fmt.Sprintf("  Fetching book names from %s...", versionName)

		resp, err := http.Get(versionLink)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("couldn't reach biblegateway version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			sp.Stop()
			msg := fmt.Sprintf("biblegateway version '%s' did not respond 200, got %d", versionName, resp.StatusCode)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}

		document, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("couldn't read biblegateway version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}

		document.Find(".book-name").Each(func(idx int, element *goquery.Selection) {
			element.Find("span").Each(func(idx int, span *goquery.Selection) {
				span.Remove()
			})

			dataName, exists := element.Attr("data-target")
			if exists {
				dataName = string([]rune(dataName)[1 : len(dataName)-5])
				bookName := strings.TrimSpace(element.Text())

				if slices.StringInSlice(dataName, []string{"3macc", "3m"}) {
					dataName = "3ma"
				} else if slices.StringInSlice(dataName, []string{"4macc", "4m"}) {
					dataName = "4ma"
				} else if slices.StringInSlice(dataName, []string{"gkesth", "adest", "addesth", "gkes"}) {
					dataName = "gkest"
				} else if slices.StringInSlice(dataName, []string{"sgthree", "sgthr", "prazar"}) {
					dataName = "praz"
				} else if dataName == "epjer" {
					return
				}

				err := isNuisance(bookName)
				if err == nil {
					if val, ok := names[dataName]; ok {
						if !slices.StringInSlice(bookName, val) {
							names[dataName] = append(names[dataName], bookName)
						}
					} else {
						names[dataName] = []string{bookName}
					}
				}
			}
		})
	}

	return names, nil
}

func getAPIBibleVersions(apiKey string, sp *spinner.Spinner) (map[string]string, error) {
	sp.Suffix = "  Fetching API.Bible versions..."

	versions := make(map[string]string)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.scripture.api.bible/v1/bibles", nil)
	if err != nil {
		sp.Stop()
		return nil, logger.LogWithError("namefetcher", "failed to create request to API.Bible", err)
	}
	req.Header.Add("api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		sp.Stop()
		return nil, logger.LogWithError("namefetcher", "couldn't reach API.Bible version list", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		sp.Stop()
		msg := fmt.Sprintf("API.Bible version list did not respond 200, got %d", resp.StatusCode)
		return nil, logger.LogWithError("namefetcher", msg, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sp.Stop()
		return nil, logger.LogWithError("namefetcher", "couldn't read API.Bible version list", err)
	}

	var abResp = new(models.ABBibleResponse)
	err = json.Unmarshal(body, &abResp)
	if err != nil {
		sp.Stop()
		return nil, logger.LogWithError("namefetcher", "failed to unmarshal API.Bible version list", err)
	}

	for _, version := range abResp.Data {
		versions[version.Name] = fmt.Sprintf("https://api.scripture.api.bible/v1/bibles/%s/books", version.ID)
	}

	return versions, nil
}

func getAPIBibleNames(versions map[string]string, apiKey string, sp *spinner.Spinner) (map[string][]string, error) {
	names := make(map[string][]string)

	client := &http.Client{}

	for versionName, versionLink := range versions {
		sp.Suffix = fmt.Sprintf("  Fetching book names from %s...", versionName)

		req, err := http.NewRequest("GET", versionLink, nil)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("failed to create request to API.Bible version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}
		req.Header.Add("api-key", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("couldn't reach API.Bible version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			sp.Stop()
			msg := fmt.Sprintf("API.Bible version '%s' did not respond 200, got %d", versionName, resp.StatusCode)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("couldn't read API.Bible version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}

		var abResp = new(models.ABBookResponse)
		err = json.Unmarshal(body, &abResp)
		if err != nil {
			sp.Stop()
			msg := fmt.Sprintf("failed to unmarshal API.Bible version '%s'", versionName)
			return nil, logger.LogWithError("namefetcher", msg, err)
		}

		for _, book := range abResp.Data {
			trueID := book.ID
			name := book.Name
			abbv := book.Abbreviation

			if len(name) == 0 || trueID == "DAG" {
				continue
			}

			name = strings.TrimSpace(name)
			id := apiBibleNames[trueID]

			if (id == "1sam" && name == "1 Kings") || (id == "2sam" && name == "2 Kings") ||
				slices.StringInSlice(abbv, []string{"3 Kings", "4 Kings"}) || (id == "lje") {
				continue
			}

			err := isNuisance(name)
			if err == nil {
				if val, ok := names[id]; ok {
					if !slices.StringInSlice(name, val) {
						names[id] = append(names[id], name)
					}
				} else {
					names[id] = []string{name}
				}
			}

			err = isNuisance(abbv)
			if err == nil {
				if val, ok := names[id]; ok {
					if !slices.StringInSlice(abbv, val) {
						names[id] = append(names[id], abbv)
					}
				} else {
					names[id] = []string{abbv}
				}
			}
		}
	}

	return names, nil
}

func isNuisance(word string) error {
	file, err := ioutil.ReadFile("./utils/namefetcher/data/nuisances.json")
	if err != nil {
		return logger.LogWithError("namefetcher", "failed to open nuisances.json", err)
	}
	json.Unmarshal(file, &nuisances)

	word = strings.ToLower(word)
	abbreviated := fmt.Sprintf("%s.", word)

	if slices.StringInSlice(word, nuisances) || slices.StringInSlice(abbreviated, nuisances) {
		return errors.New("word is nuisance")
	}

	return nil
}

type m = map[string][]string

func mergeThreeMaps(a m, b m, c m) m {
	for k := range a {
		if !slices.StringInSlice(k, defaultNames) {
			logger.LogWithError("namefetcher", fmt.Sprintf("'%s' not in default names", k), nil)
		}
	}

	for k, v := range b {
		if !slices.StringInSlice(k, defaultNames) {
			logger.LogWithError("namefetcher", fmt.Sprintf("'%s' not in default names", k), nil)
		}

		a[k] = append(a[k], v...)
	}

	for k, v := range c {
		if !slices.StringInSlice(k, defaultNames) {
			logger.LogWithError("namefetcher", fmt.Sprintf("'%s' not in default names", k), nil)
		}

		a[k] = append(a[k], v...)
	}

	return a
}
