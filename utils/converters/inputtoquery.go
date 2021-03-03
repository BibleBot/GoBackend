package converters

import (
	"encoding/json"
	"fmt"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
)

// InputToContext takes a byte slice and attempts to convert it to a Context model.
func InputToContext(input []byte) (*models.Context, error) {
	var context models.Context

	err := json.Unmarshal(input, &context)
	if err != nil {
		return nil, logger.LogWithError("inputtocontext", fmt.Sprintf("failed to convert %s to Context", string(input)), err)
	}

	return &context, nil
}
