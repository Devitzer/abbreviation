package helpers

import (
	"github.com/fatih/color"
	"fmt"
	"os"
)

func Error(msg string) {
	c := color.New(color.FgRed, color.Bold)
	errmsg := c.Sprintf("error:")

	fmt.Println(errmsg, msg)
	os.Exit(1)
}