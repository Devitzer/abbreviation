package commands

import (
	"fmt"
	"strings"
	"github.com/tidwall/gjson"
	"github.com/devitzer/abbreviation/internal/helpers"
	"github.com/fatih/color"
)

func Search(args []string, json string, userJson *string) {
	if len(args) <= 2 {
		helpers.Error("you did not name which abbreviation to search for")
	}

	var abbr string = strings.ToLower(args[2])

	meaning := gjson.Get(json, "abbreviations." + abbr)
	var userMeaning gjson.Result
	if userJson != nil {
		userMeaning = gjson.Get(*userJson, "abbreviations." + abbr)
	} else {
		// empty result to indicate it doesn't exist
		userMeaning = gjson.Result{}
	}
	// default to 1, 1: default meaning, 2: user meaning
	var use int = 1

	if !meaning.Exists() && !userMeaning.Exists() {
		helpers.Error("the abbreviation you searched for does not exist")
	}

	if !meaning.Exists() && userMeaning.Exists() {
		use = 2
	}

	bold := color.New(color.Bold)

	if use == 1 {
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
	} else if use == 2 {
		switch userMeaning.Type {
		case gjson.String:
			fmt.Println(bold.Sprintf("Meaning:"), userMeaning.String())
		default:
			helpers.Error("the user definitions file at " + helpers.GetConfigPath() + " is formatted incorrectly. Please report this issue on GitHub.")
		}
	}
}