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
	}

	absInputsToOutputs := make(map[string]string)

	for _, file := range encryptedFiles {
		localInputPath := fmt.Sprintf("./data/usx/%s.tar.zst.gpg", file)
		localOutputPath := fmt.Sprintf("./data/usx/%s.tar", file)

		absInputPath, err := filepath.Abs(localInputPath)
		if err != nil {
			logger.Log("err", "main@extractAllData", fmt.Sprintf("couldn't get absolute path of input: %s", localInputPath))
			return err
		}

		absOutputPath, err := filepath.Abs(localOutputPath)
		if err != nil {
			logger.Log("err", "main@extractAllData", fmt.Sprintf("couldn't get absolute path of output: %s", localOutputPath))
			return err
		}

		absInputsToOutputs[absInputPath] = absOutputPath
	}

	failed := false
	var failedInput string
	for input, output := range absInputsToOutputs {
		logger.Log("info", "main@extractAllData", fmt.Sprintf("extracting %s to %s", input, output))

		if extractdata.ExtractData(input, output, password) != nil {
			failed = true
			failedInput = input
		}
	}

	if failed {
		return fmt.Errorf("failed to extract: %s", failedInput)
	}

	logger.Log("info", "main@extractAllData", "extraction successful")
	return nil
}
