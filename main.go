package main

import (
    "os"
    "fmt"
    "log"
    "crypto/tls"

    "backend/logger"

    "backend/internal/commands"
    "backend/internal/verses"

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

    // Here's the fun HTTPS stuff.
    cert, err := tls.LoadX509KeyPair("https/ssl.cert", "https/ssl.key")
    if err != nil {
        logger.Log("err", "https", err.Error());
        os.Exit(1);
    }

    sslConfig := &tls.Config{Certificates: []tls.Certificate{ cert }}

    httpsListener, err := tls.Listen("tcp", ":443", sslConfig)
    if err != nil {
        logger.Log("err", "https", err.Error());
        os.Exit(2);
    }

    log.Fatal(app.Listener(httpsListener))
}
