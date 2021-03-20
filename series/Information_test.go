package series

import (
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_MemSize(t *testing.T) {
	s := NewSeries("col", 12, 43, 90, 10, 11)
	// assert that the method returns the correct number of bytes occupied
	assert.Equal(t, helpers.GetRealSizeOf(s.Data), s.MemSize())
}

func TestSeries_ValueCounts(t *testing.T) {
	s := NewSeries("col", 12, 43.1, 90, 10, 10, 10, 12, 11.5)

	// assert that the function returns the correct map
	expected := map[interface{}]interface{}{float64(12): 2, 43.1: 1, float64(90): 1, float64(10): 3, 11.5: 1}
	assert.Equal(t, expected, s.ValueCounts())
}
