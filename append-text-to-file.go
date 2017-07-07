package golangSnippets

import (
	"os"
	"strings"
)

//AppendStringToFile appends strData to filepath
//If unique is true then unique values will be inserted in file
func AppendStringToFile(strData string, filepath string, unique bool) {
	file, _ := os.OpenFile("junk.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	b := make([]byte, 1000) //this can be efficient
	file.Read(b)
	//Appending Unique
	if unique == true {
		if !strings.Contains(string(b), strData) {
			file.WriteString(strData + "\n")
		}
	} else {
		file.WriteString(strData + "\n")
	}
}

//Int & Json data later
