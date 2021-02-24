package helpers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	arrCorrect := []interface{}{2, 6, 9, 10, 5, 3}
	arrIncorrect := []interface{}{2, 6, 9, 10, "foo", 3}
	filterFunction := func(element interface{}) (bool, error) {
		val, ok := element.(int)
		if ok {
			return val%2 == 0, nil
		}
		return false, errors.New("not an int")
	}

	// assert that the function returns an error correctly
	actual, err := Filter(arrIncorrect, filterFunction)
	assert.EqualError(t, err, "not an int")
	assert.Nil(t, actual)

	// assert that the function returns a slice correctly
	actual, err = Filter(arrCorrect, filterFunction)
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{2, 6, 10}, actual)
}
