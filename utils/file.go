package utils

import "strings"

func Path(filePath string) string {
	arr := strings.Split(filePath, "/")
	rmLast := arr[:len(arr)-1]
	return strings.Join(rmLast[:], "/")
}
