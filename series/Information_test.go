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
