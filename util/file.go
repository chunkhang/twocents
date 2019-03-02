package util

import (
	"fmt"
	"path/filepath"
)

// GetFiles returns the files in the directory with the given extension
func GetFiles(directory string, extension string) (files []string, err error) {
	defer Catch(&err)

	pattern := fmt.Sprintf("%s/*%s", directory, extension)
	files, err = filepath.Glob(pattern)
	Check(err)

	return
}
