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

func TestConvertSizeToString(t *testing.T) {
	// assert that bytes are represented as bytes for small sizes
	assert.Equal(t, "143 bytes", ConvertSizeToString(143))

	// assert that kilo bytes are represented correctly
	assert.Equal(t, "12.11 KB", ConvertSizeToString(12400))

	// assert that mega bytes are represented correctly
	assert.Equal(t, "12.40 MB", ConvertSizeToString(13_000_000))

	// assert that bytes are represented as bytes for small sizes
	assert.Equal(t, "23.51 GB", ConvertSizeToString(25_240_000_000))
}

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

func TestRepeatStringSlice(t *testing.T) {
	expected := []interface{}{"foo", "foo", "foo", "foo", "foo", "foo"}
	// assert that the function correctly returns a slice with correct number of repeated strings
	assert.Equal(t, expected, RepeatIntoSlice("foo", 6))
}

func TestConvertToFloat(t *testing.T) {
	// assert that float values will be returned correctly
	v, ok := ConvertToFloat(12.43)
	assert.Equal(t, *v, 12.43)
	assert.Equal(t, ok, true)

	// assert that int values will be returned correctly after converting to float
	v, ok = ConvertToFloat(12)
	assert.Equal(t, *v, float64(12))
	assert.Equal(t, ok, true)

	// assert that function returns nil when string type is given
	v, ok = ConvertToFloat("foo")
	assert.Nil(t, v)
	assert.Equal(t, ok, false)

	// assert that function returns nil when bool type is given
	v, ok = ConvertToFloat(true)
	assert.Nil(t, v)
	assert.Equal(t, ok, false)
}
