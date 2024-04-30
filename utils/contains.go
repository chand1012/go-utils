package utils

// Contains checks if a string is in a slice of strings.
func Contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
