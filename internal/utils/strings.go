package utils

// import "sort"

// SliceContains checks if string is in slice
func SliceContains(a []string, s string) bool {
	// i := sort.SearchStrings(a, s)
	// return i < len(s) && a[i] == s
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
