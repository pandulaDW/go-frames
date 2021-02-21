package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomError(t *testing.T) {
	// assert that an error will be returned with the provided message
	msg := "this is an error"
	assert.Equal(t, errors.New(msg), CustomError(msg))
}

func TestInvalidSeriesValError(t *testing.T) {
	// assert that an error will be returned with the provided message and params
	assert.Equal(t, errors.New("invalid value at row no 4 on column test"),
		InvalidSeriesValError("foo", 4, "test"))

	// assert that blank error will be returned if the cell is blank
	assert.Equal(t, errors.New("blank value at row no 4 on column test"),
		InvalidSeriesValError("", 4, "test"))
}

func TestCustomWithStandardError(t *testing.T) {
	// assert that an error will be returned with the provided message and error string
	msg := "this is an error"
	err := errors.New("standard error message")
	assert.Equal(t, errors.New(fmt.Sprintf("%s: \n%s", msg, "standard error message")),
		CustomWithStandardError(msg, err))
}
