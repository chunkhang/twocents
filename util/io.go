package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GetFiles returns file paths in the directory with specific extension
func GetFiles(directory string, extension string) (paths []string, err error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, fmt.Sprintf(".%s", extension)) {
			path := fmt.Sprintf("%s/%s", directory, filename)
			paths = append(paths, path)
		}
	}

	return
}
