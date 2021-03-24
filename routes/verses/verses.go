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
	boProvider *providers.BollsProvider
)

// RegisterRouter registers routers related to verse processing.
func RegisterRouter(app *fiber.App, cfg *models.Config) {
	config = cfg
	abProvider = providers.NewAPIBibleProvider(cfg.APIBibleKey)
	bgProvider = providers.NewBibleGatewayProvider()
	boProvider = providers.NewBollsProvider()

	app.Get("/api/verses/fetch", fetchVerse)
}

func fetchVerse(c *fiber.Ctx) error {
	ctx, err := converters.InputToContext(c.Body(), config)
	if err != nil {
		if err == fmt.Errorf("unauth") {
			c.SendStatus(401)
		} else {
			c.SendStatus(400)
		}

		return err
	}

	str, bookSearchResults := FindBooksInString(strings.ToLower(ctx.Body))

	var response models.VerseResponse
	response.OK = true

	for _, bsr := range bookSearchResults {
		var ver models.Version

		if ctx.Prefs.Version == "" {
			if ctx.GuildPrefs.Version == "" {
				ctx.Prefs.Version = "RSV"
			} else {
				ctx.Prefs.Version = ctx.GuildPrefs.Version
			}
		}

		config.DB.Where(&models.Version{Abbreviation: ctx.Prefs.Version}).First(&ver)

		ref := GenerateReference(config.DB, str, bsr, ver)
		if ref == nil {
			continue
		}

		if !versionSupportsSection(ver, ref) {
			// TODO: make json response
			return logger.LogWithError("fetchverse", fmt.Sprintf("%s cannot be accessed from %s", ref.ToString(), ref.Version.Abbreviation), nil)
		}

		verse, err := ProcessVerse(ref, true, true)
		if err != nil {
			return err
		}

		response.Results = append(response.Results, verse)
	}

	return c.JSON(response)
}

// ProcessVerse takes a reference and formatting toggles, returning a Verse object with the result.
func ProcessVerse(ref *models.Reference, titles bool, verseNumbers bool) (*models.Verse, error) {
	var provider models.Provider

	switch ref.Version.Source {
	case "bg":
		provider = bgProvider
	case "ab":
		provider = abProvider
	default:
		return nil, logger.LogWithError("processVerse", "invalid provider found in reference", nil)
	}

	return provider.GetVerse(ref, titles, verseNumbers)
}

func versionSupportsSection(ver models.Version, ref *models.Reference) bool {
	result := true

	if (ref.IsOT && !ver.SupportsOldTestament) || (ref.IsNT && !ver.SupportsNewTestament) || (ref.IsDEU && !ver.SupportsDeuterocanon) {
		result = false
	}

	return result
}
