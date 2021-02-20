package ioread

import (
	"errors"
	"os"
)

//fileHandling will handle the file and will return the file
func fileHandling(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(errors.New("error in reading the file. " + err.Error()))
	}
	defer file.Close()

	return file
}
