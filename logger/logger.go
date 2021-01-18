package logger

// this should probably use stdlib's 'log'
// but it seems to be more complicated than what we're needing

import (
	"errors"

	"github.com/fatih/color"
)


func Log(level string, source string, msg string) error {
	if level == "err" {
		level = "erro"
	}

	prefixColor := color.New()

	switch level {
	case "info":
		prefixColor.Add(color.FgHiCyan)
	case "warn":
		prefixColor.Add(color.FgHiYellow)
	case "err":
		prefixColor.Add(color.FgHiRed)
	default:
		return errors.New("invalid log level")
	}
	
	prefixColor.Printf("[%s] ", level)

	sourceColor := color.New(color.FgHiMagenta)
	sourceColor.Printf("<%s> ", source)

	msgColor := color.New(color.FgHiGreen)
	msgColor.Println(msg)

	return nil
}