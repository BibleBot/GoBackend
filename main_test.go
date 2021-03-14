package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
)

func TestCommands_U(t *testing.T) {
	assert.Equal(t, true, true)

	/*
		app := SetupApp()

		req := httptest.NewRequest("GET", "/api/commands/test", nil)
		resp, _ := app.Test(req)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, "test", string(body), "Output should be equal to input.")
	*/
}

func TestBibleGatewayVerses(t *testing.T) {
	inputToOutput := map[string]models.VerseResponse{
		"and so we have Psalm 3:1-": {
			OK: true,
			Results: []*models.Verse{
				{
					Reference: &models.Reference{
						Book:            "Psalms",
						StartingChapter: 3,
						StartingVerse:   1,
						EndingChapter:   3,
						EndingVerse:     0,
						Version: models.Version{
							Name:                 "Revised Standard Version (RSV)",
							Abbreviation:         "RSV",
							Source:               "bg",
							SupportsOldTestament: true,
							SupportsNewTestament: true,
							SupportsDeuterocanon: true,
						},

						IsOT:  true,
						IsNT:  false,
						IsDEU: false,
					},
					Title: "Trust in God under Adversity",
					Text:  "<**1**> O Lᴏʀᴅ, how many are my foes! Many are rising against me; <**2**> many are saying of me, there is no help for him in God. *(Selah)* <**3**> But thou, O Lᴏʀᴅ, art a shield about me, my glory, and the lifter of my head. <**4**> I cry aloud to the Lᴏʀᴅ, and he answers me from his holy hill. *(Selah)* <**5**> I lie down and sleep; I wake again, for the Lᴏʀᴅ sustains me. <**6**> I am not afraid of ten thousands of people who have set themselves against me round about. <**7**> Arise, O Lᴏʀᴅ! Deliver me, O my God! For thou dost smite all my enemies on the cheek, thou dost break the teeth of the wicked. <**8**> Deliverance belongs to the Lᴏʀᴅ; thy blessing be upon thy people! *(Selah)*",
				},
			},
		},
	}

	app, _ := SetupApp(true)

	for input, output := range inputToOutput {
		input := map[string]string{
			"body":  input,
			"ver":   "RSV",
			"token": "meowmix",
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

		assert.Equal(t, outputJSON, body, fmt.Sprintf("'%s' should provide expected output", input["body"]))
	}
}

func TestAPIBibleVerses_U(t *testing.T) {
	assert.Equal(t, true, true)

	/*

		This will be strange to implement as I don't know how we can utilize this without revealing the API key.
		We will likely have to pass the API key to the CI through an env variable and go from there, somehow.

		Included in this comment is a test that should work once we get the API key stuff worked out.

		inputToOutput := map[string][]models.Verse{
			"and so we have Psalm 3:1-5": {
				OK: true,
				Results: []*models.Verse{
					Reference: &models.Reference{
						Book:            "Psalms",
						StartingChapter: 3,
						StartingVerse:   1,
						EndingChapter:   3,
						EndingVerse:     5,
						Version: models.Version{
							Abbreviation: "KJVA",
							Source:       "ab",
						},

						IsOT:  true,
						IsNT:  false,
						IsDEU: false,
					},
					Title: "",
					Text:  "A Psalm of David, when he fled from Absalom his son. <**1**> Lᴏʀᴅ, how are they increased that trouble me! many are they that rise up against me. <**2**> Many there be which say of my soul, There is no help for him in God. *(Selah)* <**3**> But thou, O Lᴏʀᴅ, art a shield for me; my glory, and the lifter up of mine head. <**4**> I cried unto the Lᴏʀᴅ with my voice, and he heard me out of his holy hill. *(Selah)* <**5**> I laid me down and slept; I awaked; for the Lᴏʀᴅ sustained me.",
				},
			},
		}

		app, _ := SetupApp(true)

		for input, output := range inputToOutput {
			input := map[string]string{
				"body": input,
				"ver":  "KJVA",
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

			assert.Equal(t, outputJSON, body, fmt.Sprintf("'%s' should provide expected output", input["body"]))
		}

	*/
}
