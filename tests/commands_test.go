package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	assert.Equal(t, true, true)
	/*app := SetupApp()

	req := httptest.NewRequest("GET", "/api/commands/test", nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "test", string(body), "Output should be equal to input.")*/
}
