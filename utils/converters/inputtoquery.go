package converters

import (
	"encoding/json"
	"fmt"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

// InputToQuery takes a byte slice and attempts to convert it to a Query model.
func InputToQuery(input []byte) (*models.Query, error) {
	var query models.Query

	err := json.Unmarshal(input, &query)
	if err != nil {
		return nil, logger.LogWithError("inputtoquery", fmt.Sprintf("failed to convert %s to Query model", string(input)), err)
	}

	return &query, nil
}