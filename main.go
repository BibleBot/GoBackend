package main

// Eventually we'll get Kamva/mgm to take care of db stuff.

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"path/filepath"

	"backend/internal/routes/commands"
	"backend/internal/routes/verses"
	"backend/internal/utils/extractdata"
	"backend/internal/utils/logger"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"
)

var version = "v1.0.0"
var fiberConfig = fiber.Config{DisableStartupMessage: true}

// Config is based off config.yml.
type Config struct {
	OwnerID       string `yaml:"ownerID"`
	APIBible      string `yaml:"apiBible"`
	DecryptionKey string `yaml:"decryptionKey"`
}

func main() {
	logger.Log("info", "init", fmt.Sprintf("BibleBot Backend %s by Evangelion Ltd.", version))

	// Create configuration from config.yml.
	config := readConfig()

	// Extract all applicable data files.
	extractAllData(config.DecryptionKey)

	// By default, we'll just serve a
	// basic HTML page indicating what's running.
	app := fiber.New(fiberConfig)
	app.Static("/", "static")

	// Cables need not apply.
	commands.RegisterRouter(app)
	verses.RegisterRouter(app)

	log.Fatal(app.ListenTLS(":443", "https/ssl.cert", "https/ssl.key"))
}

func readConfig() *Config {
	config := Config{}

	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		logger.Log("err", "config", err.Error())
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logger.Log("err", "config", err.Error())
		os.Exit(1)
	}

	return &config
}

func extractAllData(password string) error {
	encryptedFiles := []string{
		"CSB",
		"GW",
		"NASB",
		"WYC",
		"NBLA",
		"AMP",
	}

	var absInputs []string

	for _, file := range encryptedFiles {
		localInputPath := fmt.Sprintf("./data/usx/%s.tar.zst.gpg", file)

		absInputPath, err := filepath.Abs(localInputPath)
		if err != nil {
			logger.Log("err", "extractAllData", fmt.Sprintf("couldn't get absolute path of input: %s", localInputPath))
			return err
		}

		absInputs = append(absInputs, absInputPath)
	}

	failed := false
	var failedInput string
	for _, input := range absInputs {
		logger.Log("info", "extractAllData", fmt.Sprintf("extracting %s", input))

		if extractdata.ExtractData(input, password) != nil {
			failed = true
			failedInput = input
		}
	}

	if failed {
		return fmt.Errorf("failed to extract: %s", failedInput)
	}

	logger.Log("info", "extractAllData", "extraction successful")
	return nil
}
