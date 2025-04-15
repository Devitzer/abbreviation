package commands

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/fatih/color"
	"github.com/devitzer/abbreviation/internal/helpers"
)

func List(json string) {
	data := gjson.Get(json, "abbreviations")

	if !data.Exists() {
		helpers.Error("there is no abbreviations available in the abbreviation.json file.")
	}

	bold := color.New(color.Bold)

	data.ForEach(func(key, value gjson.Result) bool {
		if value.Type == gjson.String {
			fmt.Println(bold.Sprintf("%s", key.String() + ":"), bold.Sprintf("Meaning -"), value.String())
			return true
		} else if value.Type == gjson.JSON {
			if !value.IsArray() {
				helpers.Error("there is an abbreviation in the abbreviation.json file which is of an incorrect type, open a GitHub issue.")
			}

			for i, meaning := range value.Array() {
				fmt.Println(bold.Sprintf("%s", key.String() + ":"), bold.Sprintf("Meaning %v -", i + 1), meaning.String())
			}
			return true
		} else {
			helpers.Error("there is an abbreviation in the abbreviation.json file which is of an incorrect type, open a GitHub issue.")
			return false
		}
	})
}