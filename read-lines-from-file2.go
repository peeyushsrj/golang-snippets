package golangSnippets

import (
	"io/ioutil"
	"log"
	"strings"
)

//ReadLineFromFile2 will take filepath as input and return array of strings as output (without '\n')
func ReadLineFromFile2(filepath string) []string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error in reading file", err)
	}
	return strings.Split(string(b), "\n")
}
