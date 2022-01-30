package util

// SliceContains checks if string is in slice
func SliceContains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
