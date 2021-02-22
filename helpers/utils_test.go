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

func TestRange(t *testing.T) {
	// assert that the function panics when step is set as negative incorrectly
	assert.PanicsWithError(t, "step should be a negative value when high is lower than low", func() {
		Range(1, -100, 1)
	})

	// assert that the function panics when step is set as positive incorrectly
	assert.PanicsWithError(t, "step should be a positive value when high is higher than low", func() {
		Range(1, 100, -1)
	})

	// assert that the function returns a positive int slice
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, Range(1, 11, 1))

	// assert that the function returns a negative int slice
	expected = []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}
	assert.Equal(t, expected, Range(-1, -11, -1))
}
