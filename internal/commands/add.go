package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/devitzer/abbreviation/internal/helpers"
	"github.com/tidwall/sjson"
)

// this function adds the abbreviation and saves it
func addCustomAbbr(path string, key string, value string) error {
    var data []byte
    var err error

    // check if the user made a custom abbreviation before
    if _, statErr := os.Stat(path); statErr == nil {
        data, err = os.ReadFile(path)
        if err != nil {
            return err
        }
    } else {
        data = []byte(`{"abbreviations":{}}`) // start with empty structure
    }

    jsonPath := "abbreviations." + key

    // make the new json data
    updatedJSON, err := sjson.SetBytes(data, jsonPath, value)
    if err != nil {
        return err
    }

    return os.WriteFile(path, updatedJSON, 0644)
}

func Add(args []string) {
    if len(args) < 3 {
        helpers.Error("you did not give an identifier for the abbreviation, run abbreviation help add")
    }
    meaning := helpers.Prompt("Meaning of the abbreviation?: ")

    addCustomAbbr(helpers.GetConfigPath(), strings.ToLower(args[2]), meaning)
    fmt.Println("Done!")
}