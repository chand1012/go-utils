package utils

import (
	"errors"
)

// GetJSONKeys returns the keys of a JSON object as a slice of strings.
func GetJSONKeys(j any) ([]string, error) {
	var keys []string
	switch t := j.(type) {
	case map[string]interface{}:
		for key := range t {
			keys = append(keys, key)
		}
	default:
		return nil, errors.New("not a map, no keys")
	}
	return keys, nil
}
