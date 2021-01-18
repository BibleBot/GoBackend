package main

import (
	"fmt"
	"log"

	"backend/internal/routes/commands"
	"backend/internal/routes/verses"
	"backend/internal/utils/logger"

	"github.com/gofiber/fiber/v2"
)

var version = "v1.0.0"

var fiberConfig = fiber.Config{DisableStartupMessage: true}

func main() {
	logger.Log("info", "init", fmt.Sprintf("BibleBot Backend %s by Evangelion Ltd.", version))

	// By default, we'll just serve a
	// basic HTML page indicating what's running.
	app := fiber.New(fiberConfig)
	app.Static("/", "static")

	// Cables need not apply.
	commands.RegisterRouter(app)
	verses.RegisterRouter(app)

	log.Fatal(app.ListenTLS(":443", "https/ssl.cert", "https/ssl.key"))
}
