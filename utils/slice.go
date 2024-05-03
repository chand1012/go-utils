package utils

import "strings"

// StringInSlice checks if a string is in a slice of strings
func StringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// DedupeStringSlice removes duplicate strings from a slice of strings.
func DedupeStringSlice(s []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ReverseSlice reverses the order of the elements in a slice of any type.
func ReverseSlice[T any](s []T) {
	// Initialize two pointers, left and right, at the beginning and end of the slice.
	left := 0
	right := len(s) - 1

	// Swap elements at the left and right indices until they meet in the middle.
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// EndsWithAny takes a string and a slice of suffixes and returns true if the string ends with any of the suffixes.
func EndsWithAny(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// StartsWithAny takes a string and a slice of prefixes and returns true if the string starts with any of the prefixes.
func StartsWithAny(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
