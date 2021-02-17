package main

// Eventually we'll get Kamva/mgm to take care of db stuff.

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"path/filepath"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/routes/commands"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/routes/verses"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/extractdata"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/namefetcher"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"

	"golang.org/x/crypto/acme/autocert"
)

var version = "v1.0.0"
var fiberConfig = fiber.Config{DisableStartupMessage: true}

func main() {
	logger.LogInfo("init", fmt.Sprintf("BibleBot Backend %s by Kerygma Digital", version))

	app, config := SetupApp(false)

	// Set up HTTPS based on domain argument.
	var domain string
	if len(os.Args) != 2 {
		domain = "localhost"
	} else {
		domain = os.Args[1]
	}

	if domain == "localhost" {
		logger.LogInfo("init", "initialization complete. running on http://localhost")
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
			logger.LogWithError("init", "couldn't create tls listener", err)
			os.Exit(1)
		}

		logger.LogInfo("init", fmt.Sprintf("initialization complete. running at https://%s", domain))
		log.Fatal(app.Listener(ln))
	}
}

// SetupApp basically creates the normal app but without the listener and logging frills.
func SetupApp(isTest bool) (*fiber.App, *models.Config) {
	// Create configuration from config.yml.
	config := readConfig(isTest)

	// Fetch book names.
	if !isTest {
		namefetcher.FetchBookNames(config.APIBible, config.IsDryRun, false)
	}

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

	return app, config
}

func readConfig(isTest bool) *models.Config {
	var config models.Config
	path := "./config.yml"

	if isTest {
		path = "./config.test.yml"
	}

	file, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		logger.LogWithError("config", fmt.Sprintf("%s does not exist", path), err)
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
		localInputPath := fmt.Sprintf("./data/usx/%s.tar.zst.gpg", file)

		absInputPath, err := filepath.Abs(localInputPath)
		if err != nil {
			return logger.LogWithError("extractdata", fmt.Sprintf("couldn't get absolute path of input: %s", localInputPath), err)
		}

		_, err = os.Stat(absInputPath)
		if !os.IsNotExist(err) {
			absInputs = append(absInputs, absInputPath)
		} else if os.IsNotExist(err) && len(file) > 0 {
			logger.LogWarn("extractdata", fmt.Sprintf("encrypted file '%s' was specified but does not exist at %s, ignoring", file, localInputPath))
		}
	}

	failed := false
	var failedInput string
	for _, input := range absInputs {
		logger.LogInfo("extractdata", fmt.Sprintf("extracting %s", input))

		if extractdata.ExtractData(input, config.DecryptionKey) != nil {
			failed = true
			failedInput = input
		}
	}

	if failed {
		return fmt.Errorf("failed to extract: %s", failedInput)
	}

	if len(absInputs) > 0 {
		logger.LogInfo("extractdata", "extraction successful")
	}

	return nil
}
