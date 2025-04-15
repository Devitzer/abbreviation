package commands

import (
	"github.com/fatih/color"
	"github.com/devitzer/abbreviation/internal/helpers"
	"fmt"
)

func Version() {
	// make color stuff
	c := color.New(color.FgBlue, color.Bold)
	msg := c.Sprintf("abbreviation")

	fmt.Println(msg, helpers.Version())
}