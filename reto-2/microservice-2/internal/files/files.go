package files

import (
	"os"
	"strings"
)

func ListFiles(directory string) []string {
	var files []string

	dirEntries, _ := os.ReadDir(directory)
	for _, file := range dirEntries {
		files = append(files, file.Name())
	}
	return files
}

func SearchFiles(directory, name string) []string {
	files := ListFiles(directory)
	var ret []string
	if name == "*" {
		return files
	}
	for _, file := range files {
		if strings.Contains(file, name) {
			ret = append(ret, file)
		}
	}
	return ret
}
