package helpers

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Prompts the user with a question and returns the input
func Prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

	return input
}