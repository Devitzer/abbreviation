package commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/devitzer/abbreviation/internal/helpers"
)

func Help(args []string) {
	c := color.New(color.FgBlue, color.Bold)
	optional := color.New(color.FgYellow)
	required := color.New(color.FgRed)
	bold := color.New(color.Bold)

	if len(args) >= 3 {
		switch args[2] {
		case "help":
			bold.Println("Help Command")
			fmt.Println("The help command provides a list of commands or specific information on a single command.")
			fmt.Println("	Arguments:")
			fmt.Printf("	%s: The command to provide information on. (%s)\n", bold.Sprintf("command"), optional.Sprintf("OPTIONAL"))
		case "search":
			bold.Println("Search Command")
			fmt.Println("This command searches an abbreviation from the database and returns the meaning(s).")
			fmt.Println("	Arguments:")
			fmt.Printf("	%s: The abbreviation to search. (%s)\n", bold.Sprintf("abbreviation"), required.Sprintf("REQUIRED"))
		case "version":
			bold.Println("Version Command")
			fmt.Println("This command tells you the current version of the CLI that you have installed.")
		case "list":
			bold.Println("List Command")
			fmt.Println("This command lists all the abbreviations in the database.")
		default:
			var err string = args[2] + " is not a command! Type \"abbreviation help\" to get a list of commands."
			helpers.Error(err)
		}
	} else {
		fmt.Println(c.Sprintf("abbreviation |"), "List of Commands")
		fmt.Println("help")
		fmt.Println("search")
		fmt.Println("version")
		fmt.Println("list")
	}
}