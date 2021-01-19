package main

// Eventually we'll get Kamva/mgm to take care of db stuff.

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"path/filepath"

	"github.com/BibleBot/backend/internal/routes/commands"
	"github.com/BibleBot/backend/internal/routes/verses"
	"github.com/BibleBot/backend/internal/utils/extractdata"
	"github.com/BibleBot/backend/internal/utils/logger"
	"github.com/BibleBot/backend/internal/utils/namefetcher"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"

	"golang.org/x/crypto/acme/autocert"
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
	IsDryRun         bool     `yaml:"isDryRun"`
}

func main() {
	logger.Log("info", "init", fmt.Sprintf("BibleBot Backend %s by Evangelion Ltd.", version))

	// Create configuration from config.yml.
	config := readConfig()

	// Fetch book names.
	namefetcher.FetchBookNames(config.APIBible, config.IsDryRun)

	// Extract all applicable data files.
	extractAllData(config)

	// By default, we'll just serve a
	// basic HTML page indicating what's running.
	app := fiber.New(fiberConfig)
	app.Static("/", "static")

	// Cables need not apply.
	commands.RegisterRouter(app)
	verses.RegisterRouter(app)

	// Set up HTTPS based on domain argument.
	var domain string
	if len(os.Args) != 2 {
		domain = "localhost"
	} else {
		domain = os.Args[1]
	}

	if domain == "localhost" {
		log.Fatal(app.Listen(":80"))
	} else {
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Email:      config.LetsEncryptEmail,
			HostPolicy: autocert.HostWhitelist(os.Args[1]),
			Cache:      autocert.DirCache("./https"),
		}

		cfg := &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos: []string{
				"http/1.1", "acme-tls/1",
			},
		}

		ln, err := tls.Listen("tcp", ":443", cfg)
		if err != nil {
			logger.Log("err", "init", "couldn't create tls listener")
			os.Exit(1)
		}

		log.Fatal(app.Listener(ln))
	}
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

func extractAllData(config *Config) error {
	var absInputs []string

	for _, file := range config.EncryptedFiles {
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

		if extractdata.ExtractData(input, config.DecryptionKey) != nil {
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
