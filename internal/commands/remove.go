package commands

import (
	"os"
    "fmt"
    "strings"

	"github.com/tidwall/sjson"
    "github.com/tidwall/gjson"
    "github.com/devitzer/abbreviation/internal/helpers"
)

// this function adds the abbreviation and saves it
func removeCustomAbbr(path string, key string) error {
    var data []byte
    var err error

    // check if the user made a custom abbreviation before
    if _, statErr := os.Stat(path); statErr == nil {
        data, err = os.ReadFile(path)
        if err != nil {
            return err
        }
    } else {
        helpers.Error("there is no abbreviations to delete because no custom ones have been made")
    }

    jsonPath := "abbreviations." + key

    if !gjson.GetBytes(data, jsonPath).Exists() {
        helpers.Error("you cannot delete that abbreviation because it does not exist")
    }

    // Delete the key
    updatedJSON, err := sjson.DeleteBytes(data, jsonPath)
    if err != nil {
        return err
    }

    return os.WriteFile(path, updatedJSON, 0644)
}

func Remove(args []string) {
    if len(args) < 3 {
        helpers.Error("you did not give an identifier for the abbreviation, run abbreviation help remove")
    }

    removeCustomAbbr(helpers.GetConfigPath(), strings.ToLower(args[2]))
    fmt.Println("The abbreviation was deleted.")
}