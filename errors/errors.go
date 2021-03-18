package errors

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
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

// DuplicatedColumn returns an error to indicate the specified column is already in the dataframe
func DuplicatedColumn(column string) error {
	return errors.New(fmt.Sprintf("%s column is already in the dataframe", column))
}

// MismatchedNumOfColumns provides an error mentioning the mismatched number of columns
func MismatchedNumOfColumns(expected, actual int) error {
	err := fmt.Sprintf("mismatched number of columns provided. requires %d columns, but %d was provided",
		expected, actual)
	return errors.New(err)
}

// MismatchedNumOfRows provides an error mentioning the mismatched number of columns
func MismatchedNumOfRows(expected, actual int) error {
	err := fmt.Sprintf("mismatched number of rows provided. requires %d rows, but %d was provided",
		expected, actual)
	return errors.New(err)
}

// IncorrectDataType returns an error mentioning the expected type
func IncorrectDataType(dtype base.DType) error {
	return errors.New(fmt.Sprintf("expected a %s type Series", dtype))
}

// InvalidRowValue returns an error if an incorrect value is found in a series at the given row number
func InvalidRowValue(rowNum int) error {
	return errors.New(fmt.Sprintf("invalid value at row %d", rowNum))
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
