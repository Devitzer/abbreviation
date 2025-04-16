package helpers

// this command returns a directory to save custom abbreviations depending on the user's OS

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/tidwall/gjson"
)

func GetConfigPath() string {
	var configDir string
	// find user's home dir (linux = ~/, windows = C:/User/<user here>/)
	home, err := os.UserHomeDir()
    if err != nil {
		Error("could not find your home dir, are you on a platform other than Windows, macOS or Linux? Your platform may not support custom abbreviations.")
    }

	if runtime.GOOS == "windows" {
        configDir = filepath.Join(home, "AppData", "Local", "abbreviation")
    } else {
        configDir = filepath.Join(home, ".config", "abbreviation")
    }

	os.MkdirAll(configDir, os.ModePerm)
    return filepath.Join(configDir, "custom_abbr.json")
}

func FileExists(path string) bool {
    info, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil && !info.IsDir()
}

func LoadCustomAbbr(path string) gjson.Result {
    data, err := os.ReadFile(path)
    if err != nil {
        return gjson.Result{} // empty result if file missing/unreadable
    }

    return gjson.GetBytes(data, "abbreviations")
}
