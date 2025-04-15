package main

import (
	"github.com/devitzer/abbreviation/internal/commands"
	"github.com/devitzer/abbreviation/internal/helpers"
	"os"
	_ "embed"
)

//go:embed abbreviation.json
var abbreviations string

func main() {
	args := os.Args

	if len(args) <= 1 {
		helpers.Error("you did not enter a command to execute")
	}

	cmd := args[1]

	// argument parsing is handled within the commands if they involve arguments
	switch cmd {
		case "version":
			commands.Version()
		case "help":
			commands.Help(args)
		case "search":
			commands.Search(args, abbreviations)
		case "list":
			commands.List(abbreviations)
	}
}