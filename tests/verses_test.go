package tests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerses(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("GET", "/api/verses/test", nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "test", string(body), "Output should be equal to input.")
}
