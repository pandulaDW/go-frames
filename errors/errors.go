package errors

import (
	"errors"
	"fmt"
)

// CustomError will return a custom error based on the message provided
func CustomError(msg string) error {
	return errors.New(msg)
}

// CustomWithStandardError will return a custom error massage combined with a standard error message
func CustomWithStandardError(msg string, err error) error {
	return errors.New(fmt.Sprintf("%s: \n%s", msg, err.Error()))
}

// ColumnNotFound returns an error to indicate the specified column is not found in the dataframe
func ColumnNotFound(column string) error {
	return errors.New(fmt.Sprintf("%s column not found in the dataframe", column))
}

// ErrMismatchedNumOfColumns provides an error mentioning the mismatched number of columns
func MismatchedNumOfColumns(expected, actual int) error {
	err := fmt.Sprintf("mismatched number of columns provided. requires %d columns, but %d was provided",
		expected, actual)
	return errors.New(err)
}

//InvalidSeriesValError return an error based an error specifying the index and column name
func InvalidSeriesValError(val interface{}, i int, col string) error {
	var errStr string
	if val == "" {
		errStr = fmt.Sprintf("blank value at row no %d on column %s", i, col)
	} else {
		errStr = fmt.Sprintf("invalid value at row no %d on column %s", i, col)
	}
	return errors.New(errStr)
}
