package verses

// Parse message with parsing.go, then generate a reference here and process it.

import (
	"encoding/json"
	"fmt"

	"github.com/BibleBot/backend/internal/models"
	"github.com/BibleBot/backend/internal/utils/converters"
	"github.com/BibleBot/backend/internal/utils/logger"
	"github.com/gofiber/fiber/v2"
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

	str, results := FindBooksInString(query.Body)
	result, err := json.Marshal(results)
	if err != nil {
		return logger.LogWithError("verses@fetchVerse", "unable to convert result to json", nil)
	}
	fmt.Println(str)
	fmt.Println(string(result))
	return c.SendString(string(result))
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
