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
