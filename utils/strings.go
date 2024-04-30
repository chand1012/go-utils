package utils

import "regexp"

// RemoveQuotes removes any quotes from a string using regex.
func RemoveQuotes(s string) string {
	// regex to match any quotes
	regex := `["']`
	// replace any quotes with an empty string
	return regexp.MustCompile(regex).ReplaceAllString(s, "")
}
