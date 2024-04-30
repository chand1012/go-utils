package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input prompts the user for input and returns the entered text.
func Input(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// InputWithColor prompts the user for input and returns the entered text. The prompt is displayed in the specified CSS color code.
func InputWithColor(prompt, cssColorCode string) (string, error) {
	PrintColoredText(prompt, cssColorCode)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}
