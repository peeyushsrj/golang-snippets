package golangSnippets

import (
	"io/ioutil"
	"log"
	"strings"
)

//ReadLineFromFile2 will take filepath as input and return array of strings as output (without '\n')
func ReadLineFromFile2(filepath string) ([]string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		var emptyArr []string
		return emptyArr, err 
	}
	return strings.Split(string(b), "\n"), nil
}
