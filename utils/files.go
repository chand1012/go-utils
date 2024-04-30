package utils

import (
	"errors"
	"os"
	"unicode/utf8"
)

// WriteFile writes a string to a file at the specified path.
func WriteFile(filePath string, contents string) error {
	return os.WriteFile(filePath, []byte(contents), 0644)
}

// LoadFile reads the contents of a file at the specified path and returns it as a string. An error is returned if the file cannot be read or if it is not valid UTF-8.
func LoadFile(filePath string) (string, error) {
	// load the file
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	if !utf8.Valid(contents) {
		return "", errors.New("the file is not valid UTF-8")
	}

	return string(contents), nil
}
