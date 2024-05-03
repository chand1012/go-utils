package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// RemoveQuotes removes any quotes from a string using regex.
func RemoveQuotes(s string) string {
	// regex to match any quotes
	regex := `["']`
	// replace any quotes with an empty string
	return regexp.MustCompile(regex).ReplaceAllString(s, "")
}

// InsertLine inserts a new line of text at the specified line number in the code.
func InsertLine(code string, lineNumber int, newText string) (string, error) {
	if lineNumber < 1 {
		return "", fmt.Errorf("line number must be greater than or equal to 1")
	}

	lines := strings.Split(code, "\n")
	if lineNumber > len(lines)+1 {
		return "", fmt.Errorf("line number is greater than the number of lines in the code")
	}

	lines = append(lines[:lineNumber-1], append([]string{newText}, lines[lineNumber-1:]...)...)
	return strings.Join(lines, "\n"), nil
}

// InsertLinesAtIndices inserts multiple lines of text at the specified line numbers in the code.
func InsertLinesAtIndices(file string, indices []int, linesToInsert []string) (string, error) {
	if len(indices) != len(linesToInsert) {
		return "", fmt.Errorf("the length of indices and linesToInsert must be the same")
	}

	lines := strings.Split(file, "\n")

	for i, index := range indices {
		if index < 1 || index > len(lines)+1 {
			return "", fmt.Errorf("index %d is out of bounds", index)
		}

		lines = append(lines[:index-1+i], append([]string{linesToInsert[i]}, lines[index-1+i:]...)...)
	}

	return strings.Join(lines, "\n"), nil
}

// ReplaceLines replaces the lines between the startLine and endLine with the new text.
func ReplaceLines(code string, startLine int, endLine int, newText string) (string, error) {
	if startLine < 1 || endLine < 1 {
		return "", fmt.Errorf("startLine and endLine must be greater than or equal to 1")
	}

	if startLine > endLine {
		return "", fmt.Errorf("startLine must be less than or equal to endLine")
	}

	lines := strings.Split(code, "\n")
	if endLine > len(lines) {
		return "", fmt.Errorf("endLine is greater than the number of lines in the code")
	}

	lines = append(lines[:startLine-1], append([]string{newText}, lines[endLine:]...)...)
	return strings.Join(lines, "\n"), nil
}
