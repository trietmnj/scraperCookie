package util

import "strings"

// Path remove last element from a filePath as delineated with /
func Path(filePath string) string {
	arr := strings.Split(filePath, "/")
	rmLast := arr[:len(arr)-1]
	return strings.Join(rmLast[:], "/")
}

// File returns the last element from a  filePath as delineated with /
func File(filePath string) string {
	arr := strings.Split(filePath, "/")
	rmLast := arr[len(arr)]
	return rmLast
}
