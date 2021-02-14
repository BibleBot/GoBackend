package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

func TestVerses(t *testing.T) {
	inputToOutput := map[string][]models.Verse{
		"and so we have Psalm 3:1-": {
			{
				Reference: &models.Reference{
					Book:            "Psalms",
					StartingChapter: 3,
					StartingVerse:   1,
					EndingChapter:   3,
					EndingVerse:     0,
					Version: models.Version{
						Abbreviation: "RSV",
					},

					IsOT:  true,
					IsNT:  false,
					IsDEU: false,
				},
				Title: "Trust in God under Adversity",
				Text:  "<**1**> O Lord, how many are my foes! Many are rising against me; <**2**> many are saying of me, there is no help for him in God. *(Selah)* <**3**> But thou, O Lord, art a shield about me, my glory, and the lifter of my head. <**4**> I cry aloud to the Lord, and he answers me from his holy hill. *(Selah)* <**5**> I lie down and sleep; I wake again, for the Lord sustains me. <**6**> I am not afraid of ten thousands of people who have set themselves against me round about. <**7**> Arise, O Lord! Deliver me, O my God! For thou dost smite all my enemies on the cheek, thou dost break the teeth of the wicked. <**8**> Deliverance belongs to the Lord; thy blessing be upon thy people! *(Selah)*",
			},
		},
	}

	app := SetupApp()

	for input, output := range inputToOutput {
		input := map[string]string{
			"body": input,
		}

		b, err := json.Marshal(input)
		if err != nil {
			log.Fatal(err)
		}

		req := httptest.NewRequest("GET", "/api/verses/fetch", bytes.NewBuffer(b))
		resp, _ := app.Test(req)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		outputJSON, err := json.Marshal(output)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, outputJSON, body, "Requesting Psalm 3:1- should give Psalm 3:1 and the rest of the chapter.")
	}
}
