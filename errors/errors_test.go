package errors

import (
	"errors"
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
