package golangSnippets

import "strings"

//stringInSlice finds a string inside an array of strings, 
//if found return the string else return empty string
func StringInSlice(a string, b []string) string {
	for _, el := range b {
		if strings.Contains(a, el) {
			return el
		}
	}
	return ""
}
