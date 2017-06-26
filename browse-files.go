package golangSnippets

import (
	"os"
	"path/filepath"
	"strings"
)

//Browse files of type x (e.g. `.mp3`) from root directory path, which returns array of file paths and error if exist.
func BrowseXFiles(x string, root string) ([]string, error) {
	var arr []string
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			if strings.HasSuffix(f.Name(), x) { //.mp3
				arr = append(arr, path)
			}
		}
		return nil
	})
	if err != nil {
		return arr, err
	}
	return arr, nil
}
