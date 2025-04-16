package main

import (
	"github.com/devitzer/abbreviation/internal/commands"
	"github.com/devitzer/abbreviation/internal/helpers"
	"os"
	_ "embed"
)

//go:embed abbreviation.json
var abbreviations string

var userDefinitionBytes []byte
var userDefinitions *string

func main() {
	args := os.Args

	if len(args) <= 1 {
		helpers.Error("you did not enter a command to execute, run abbreviation help for a list of commands")
	}

	if helpers.FileExists(helpers.GetConfigPath()) {
		userDefinitionBytes, _ = os.ReadFile(helpers.GetConfigPath())
		str := string(userDefinitionBytes)
		userDefinitions = &str
	} else {
		userDefinitions = nil
	}

	cmd := args[1]

	// argument parsing is handled within the commands if they involve arguments
	switch cmd {
		case "version":
			commands.Version()
		case "help":
			commands.Help(args)
		case "search":
			commands.Search(args, abbreviations, userDefinitions)
		case "list":
			commands.List(abbreviations)
		case "add":
			commands.Add(args)
		case "remove":
			commands.Remove(args)
		case "list-user":
			commands.ListUser(userDefinitions)
	}
}