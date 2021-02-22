package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueCounts(t *testing.T) {
	arr := []interface{}{"foo", "bar", "bar", "baz", "foo", "foo", "dar"}
	expected := map[interface{}]int{"foo": 3, "bar": 2, "baz": 1, "dar": 1}

	// assert that the function returns correct value counts
	assert.Equal(t, expected, ValueCounts(arr))
}

func TestMaxIntSlice(t *testing.T) {
	arr := []int{12, 45, 65, 56, 90, 81, 22}
	// assert that the function returns the correct max value
	assert.Equal(t, 90, MaxIntSlice(arr))
}

func TestReverseArray(t *testing.T) {
	arr := []interface{}{12, 45, 65, 56}
	// assert that the function reverses the array correctly
	assert.Equal(t, []interface{}{56, 65, 45, 12}, ReverseArray(arr))
}
