package series

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//goland:noinspection GoNilness
func TestSeries_Round(t *testing.T) {
	s := NewSeries("col", 12.346, 56.23, 1.2, 1.2335, 4.6003, 1.34)
	expected := NewSeries("round(col)", 12.35, 56.23, 1.20, 1.23, 4.60, 1.34)

	// assert that the function will return a rounded series
	assert.Equal(t, s.Round(2), expected)

	// assert that the function will panic if a wrong type series is given
	assert.PanicsWithError(t, "expected a Float type Series", func() {
		NewSeries("col", "foo", "bar").Round(2)
	})

	// assert that the function will panic if there's an invalid value
	s.Data[2] = "foo"
	assert.PanicsWithError(t, fmt.Sprintf("invalid value at row %d", 2), func() {
		s.Round(2)
	})
}
