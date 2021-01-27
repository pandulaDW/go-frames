package errors

import (
	"errors"
	"fmt"
)

// CustomError will return a custom error based on the message provided
func CustomError(msg string) error {
	return errors.New(msg)
}

//InvalidSeriesValError return an error based an error specifying the index and column name
func InvalidSeriesValError(i int, col string) error {
	errStr := fmt.Sprintf("invalid value at row no %d on column %s", i, col)
	return errors.New(errStr)
}