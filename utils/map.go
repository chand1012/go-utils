package utils

// BreakOutMap takes a map and returns two slices, one containing the keys and the other containing the values.
func BreakOutMap(m map[string]string) (keys []string, values []string) {
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
