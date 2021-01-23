package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRealSizeOf(t *testing.T) {
	intSlice := []byte{2, 3, 5, 6, 9}

	// bytes occupied by the intSlice (hard to approximate)
	expected := 9

	// assert that function returns the correct number of bytes occupied by an interface
	assert.Equal(t, expected, GetRealSizeOf(intSlice))

	// assert that function panics when supplied with a nil pointer
	var nilVal interface{}
	assert.Panics(t, func() {
		GetRealSizeOf(nilVal)
	})
}

func TestValueCounts(t *testing.T) {
	arr := []interface{}{"foo", "bar", "bar", "baz", "foo", "foo", "dar"}
	expected := map[interface{}]int{"foo": 3, "bar": 2, "baz": 1, "dar": 1}

	// assert that the function returns correct value counts
	assert.Equal(t, expected, ValueCounts(arr), "returns a correct value count map")
}

func TestMaxIntSlice(t *testing.T) {
	arr := []int{12, 45, 65, 56, 90, 81, 22}
	assert.Equal(t, 90, MaxIntSlice(arr), "returns the correct max value")
}
