package series

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_Round(t *testing.T) {
	s := NewSeries("col", 12.346, 56.23, 1.2, 1.2335, 4.6003, 1.34)
	expected := NewSeries("col", 12.35, 56.23, 1.20, 1.23, 4.60, 1.34)

	// assert that the function will mutate the the series data correctly if inplace is true
	assert.Equal(t, s.Round(2, true), expected)
}
