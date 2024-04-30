package utils

// StringSliceToAnySlice converts a slice of strings to a slice of any.
func StringSliceToAnySlice(s []string) []any {
	var i []any
	for _, v := range s {
		i = append(i, v)
	}
	return i
}
