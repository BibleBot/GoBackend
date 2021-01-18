package main

import (
	"fmt"
	"log"

	"backend/logger"

	"backend/internal/commands"
	"backend/internal/verses"

	"github.com/gofiber/fiber/v2"
)

var version = "v1.0.0"

func main() {
	logger.Log("info", "global", fmt.Sprintf("BibleBot Backend %s by Evangelion Ltd.", version))

	app := fiber.New()

	app.Static("/", "static")

	commands.RegisterRouter(app)
	verses.RegisterRouter(app)

	log.Fatal(app.Listen(":3000"))
}
