package golangSnippets

import (
	"strings"
)

//stringInSlice finds a string inside an array of strings
func stringInSlice(a string, b []string) bool {
	for _, el := range b {
		if strings.Contains(a, el) {
			return true
		}
	}
	return false
}
