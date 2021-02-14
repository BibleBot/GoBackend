package verses

// Parse message with parsing.go, then generate a reference here and process it.

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/routes/verses/interfaces"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/converters"
)

var _config *models.Config

// RegisterRouter registers routers related to verse processing.
func RegisterRouter(app *fiber.App, config *models.Config) {
	_config = config

	app.Get("/api/verses/fetch", fetchVerse)
}

func fetchVerse(c *fiber.Ctx) error {
	query, err := processInput(c.Body())
	if err != nil {
		if err == fmt.Errorf("unauth") {
			c.SendStatus(401)
		} else {
			c.SendStatus(400)
		}

		return err
	}

	str, bookSearchResults := FindBooksInString(strings.ToLower(query.Body))

	var verseResults []*models.Verse

	for _, bsr := range bookSearchResults {
		reference := GenerateReference(str, bsr, models.Version{
			Abbreviation: "RSV",
		})
		if reference == nil {
			continue
		}

		verse, err := interfaces.GetBibleGatewayVerse(reference, true, true)
		if err != nil {
			return err
		}

		verseResults = append(verseResults, verse)
	}

	return c.JSON(verseResults)
}

func processInput(input []byte) (*models.Query, error) {
	query, err := converters.InputToQuery(input)
	if err != nil {
		return nil, err
	}

	//if query.Token != _config.AccessKey {
	//	return nil, errors.New("unauth")
	//}

	return query, nil
}
