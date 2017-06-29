package golangSnippets

import (
	"log"
	"net/url"
	"path/filepath"
)

func FilenameFromUrl(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal("Error due to parsing url: ", err)
	}
	return filepath.Base(u.EscapedPath())
}
