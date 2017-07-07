package golangSnippets

import (
	"os"
	"strings"
)

//AppendToFile appends data to filepath named file
//If unique is true then unique values will be inserted in file
func AppendToFile(data interface{}, filepath string, unique bool) {
	fi, _ := os.OpenFile("junk.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer fi.Close()
	b := make([]byte, 1000) //this can be efficient
	file.Read(b)
	//Appending Unique
	if uniqueFlag == true {
		if !strings.Contains(string(b), data) {
			file.WriteString(data + "\n")
		}
	} else {
		file.WriteString(data + "\n")
	}
}
