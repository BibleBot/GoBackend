package logger

// this should probably use stdlib's 'log'
// but it seems to be more complicated than what we're needing

import (
	"errors"

	"github.com/fatih/color"
)

func log(level string, source string, msg string) error {
	if level == "err" {
		level = "erro"
	}

	prefixColor := color.New()

	switch level {
	case "info":
		prefixColor.Add(color.FgHiCyan)
	case "warn":
		prefixColor.Add(color.FgHiYellow)
	case "erro":
		prefixColor.Add(color.FgHiRed)
	default:
		return errors.New("invalid log level")
	}

	prefixColor.Printf("[%s] ", level)

	sourceColor := color.New(color.FgHiMagenta)
	sourceColor.Printf("<%s> ", source)

	msgColor := color.New(color.Reset)

	if source == "init" {
		msgColor.Add(color.FgHiGreen)
	}

	msgColor.Println(msg)

	return nil
}

// LogInfo wraps log() in reporting a standard informational message.
func LogInfo(source string, message string) error {
	return log("info", source, message)
}

// LogWarn wraps log() in reporting a standard warning message.
func LogWarn(source string, message string) error {
	return log("warn", source, message)
}

// LogWithError wraps log() in reporting a standard error message and returning the error.
func LogWithError(source string, msg string, err error) error {
	log("err", source, msg)

	if err != nil {
		return errors.New(msg)
	}

	return err
}
