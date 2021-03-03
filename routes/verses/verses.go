package verses

// Parse message with parsing.go, then generate a reference here and process it.

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/routes/verses/providers"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/converters"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

var (
	config     *models.Config
	abProvider *providers.APIBibleProvider
	bgProvider *providers.BibleGatewayProvider
)

// RegisterRouter registers routers related to verse processing.
func RegisterRouter(app *fiber.App, cfg *models.Config) {
	config = cfg
	abProvider = providers.NewAPIBibleProvider(cfg.APIBibleKey)
	bgProvider = providers.NewBibleGatewayProvider()

	app.Get("/api/verses/fetch", fetchVerse)
}

func fetchVerse(c *fiber.Ctx) error {
	ctx, err := processInput(c.Body())
	if err != nil {
		if err == fmt.Errorf("unauth") {
			c.SendStatus(401)
		} else {
			c.SendStatus(400)
		}

		return err
	}

	str, bookSearchResults := FindBooksInString(strings.ToLower(ctx.Body))

	var verseResults []*models.Verse

	for _, bsr := range bookSearchResults {
		rsv := models.Version{
			Abbreviation: "RSV",
			Source:       "bg",
		}

		kjva := models.Version{
			Abbreviation: "KJVA",
			Source:       "ab",
		}

		var ver models.Version

		switch ctx.TempVersion {
		case "RSV":
			ver = rsv
			break
		case "KJVA":
			ver = kjva
			break
		}

		reference := GenerateReference(str, bsr, ver)
		if reference == nil {
			continue
		}

		verse, err := ProcessVerse(reference, true, true)
		if err != nil {
			return err
		}

		verseResults = append(verseResults, verse)
	}

	return c.JSON(verseResults)
}

// ProcessVerse takes a reference and formatting toggles, returning a Verse object with the result.
func ProcessVerse(ref *models.Reference, titles bool, verseNumbers bool) (*models.Verse, error) {
	var provider models.Provider

	switch ref.Version.Source {
	case "bg":
		provider = bgProvider
		break
	case "ab":
		provider = abProvider
		break
	default:
		return nil, logger.LogWithError("processVerse", "invalid provider found in reference", nil)
	}

	return provider.GetVerse(ref, titles, verseNumbers)
}

func processInput(input []byte) (*models.Context, error) {
	context, err := converters.InputToContext(input)
	if err != nil {
		return nil, err
	}

	//if query.Token != _config.AccessKey {
	//	return nil, errors.New("unauth")
	//}

	return context, nil
}
