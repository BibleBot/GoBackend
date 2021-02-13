package tests

// Eventually we'll get Kamva/mgm to take care of db stuff.

import (
	"fmt"
	"io/ioutil"
	"os"

	"path/filepath"

	"github.com/BibleBot/backend/internal/models"
	"github.com/BibleBot/backend/internal/routes/commands"
	"github.com/BibleBot/backend/internal/routes/verses"
	"github.com/BibleBot/backend/internal/utils/extractdata"
	"github.com/BibleBot/backend/internal/utils/logger"
	"github.com/BibleBot/backend/internal/utils/namefetcher"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"
)

var version = "v1.0.0"
var fiberConfig = fiber.Config{DisableStartupMessage: true}

// Config is based off config.yml.
type Config struct {
	OwnerID          string   `yaml:"ownerID"`
	APIBible         string   `yaml:"apiBible"`
	LetsEncryptEmail string   `yaml:"letsEncryptEmail"`
	DecryptionKey    string   `yaml:"decryptionKey"`
	EncryptedFiles   []string `yaml:"encryptedFiles"`
	IsDryRun         bool     `yaml:"dry"`
}

// SetupApp basically creates the normal app but without the listener and logging frills.
func SetupApp() *fiber.App {
	// Create configuration from config.yml.
	config := readConfig()

	// Fetch book names.
	namefetcher.FetchBookNames(config.APIBible, config.IsDryRun, true)

	// Extract all applicable data files.
	err := extractData(config)
	if err != nil {
		os.Exit(1)
	}

	// By default, we'll just serve a basic HTML page indicating what's running.
	app := fiber.New(fiberConfig)
	app.Static("/", "static")

	// Cables need not apply.
	commands.RegisterRouter(app)
	verses.RegisterRouter(app, config)

	return app
}

func readConfig() *models.Config {
	var config models.Config

	file, err := ioutil.ReadFile("./config.test.yml")
	if os.IsNotExist(err) {
		logger.LogWithError("config", "config.test.yml does not exist", err)
		os.Exit(1)
	} else if err != nil {
		logger.LogWithError("config", err.Error(), err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logger.LogWithError("config", err.Error(), err)
		os.Exit(1)
	}

	return &config
}

// This is a wrapper for internal/utils/extractdata.
// We're using the same name just to make logging more uniform.
func extractData(config *models.Config) error {
	var absInputs []string

	for _, file := range config.EncryptedFiles {
		localInputPath := fmt.Sprintf("./../data/usx/%s.tar.zst.gpg", file)

		absInputPath, err := filepath.Abs(localInputPath)
		if err != nil {
			return err
		}

		_, err = os.Stat(absInputPath)
		if !os.IsNotExist(err) {
			absInputs = append(absInputs, absInputPath)
		}
	}

	failed := false
	var failedInput string
	for _, input := range absInputs {
		if extractdata.ExtractData(input, config.DecryptionKey) != nil {
			failed = true
			failedInput = input
		}
	}

	if failed {
		return fmt.Errorf("failed to extract: %s", failedInput)
	}

	return nil
}
