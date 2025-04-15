package commands

import (
	"fmt"
	"strings"
	"github.com/tidwall/gjson"
	"github.com/devitzer/abbreviation/internal/helpers"
	"github.com/fatih/color"
)

func Search(args []string, json string) {
	if len(args) <= 2 {
		helpers.Error("you did not name which abbreviation to search for")
	}

	var abbr string = strings.ToLower(args[2])

	meaning := gjson.Get(json, "abbreviations." + abbr)

	if !meaning.Exists() {
		helpers.Error("the abbreviation you searched for does not exist!")
	}

	bold := color.New(color.Bold)

	switch meaning.Type {
	case gjson.String:
		fmt.Println(bold.Sprintf("Meaning:"), meaning.String())
	case gjson.JSON:
		if !meaning.IsArray() {
			helpers.Error("the abbreviation.json file is formatted incorrectly. Please report this issue on GitHub.")
		}
		for i, mean := range meaning.Array() {
			fmt.Println(bold.Sprintf("Meaning %v:", i + 1), mean.String())
		}
	default:
		helpers.Error("the abbreviation.json file is formatted incorrectly. Please report this issue on GitHub.")
	}
}