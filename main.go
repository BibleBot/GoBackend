package main

import (
	"fmt"

	"./internal/commands"
	"./internal/verses"
)

func main() {
	verses.Router()
	commands.Router()
	fmt.Printf("Hello world.")
}
