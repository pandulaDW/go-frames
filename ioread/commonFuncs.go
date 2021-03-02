package ioread

import (
	"github.com/pandulaDW/go-frames/errors"
	"os"
)

// fileHandling will handle the file and will return the file
func fileHandling(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.CustomWithStandardError("error in reading the file", err)
	}

	return file, nil
}
