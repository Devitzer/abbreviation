package commands

import (
	"github.com/fatih/color"
	"github.com/devitzer/abbreviation/internal/helpers"
	"fmt"
	"runtime"
)

func Version() {
	// make color stuff
	c := color.New(color.FgBlue, color.Bold)
	bold := color.New(color.Bold)
	msg := c.Sprintf("abbreviation")

	fmt.Printf("%s %s %s\n", msg, bold.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), helpers.Version())
}