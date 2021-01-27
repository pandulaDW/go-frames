package series

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

//goland:noinspection GoNilness
func TestSeries_Round(t *testing.T) {
	s := NewSeries("col", 12.346, 56.23, 1.2, 1.2335, 4.6003, 1.34)
	expected := NewSeries("col", 12.35, 56.23, 1.20, 1.23, 4.60, 1.34)

	// assert that the function will mutate the the series data correctly if inplace is true
	assert.Equal(t, s.Copy().Round(2, true), expected)

	// assert that the function will return a new series if inplace is false
	assert.Equal(t, s.Round(2, false), expected)
	assert.NotEqual(t, s.Round(2, false), s)

	// assert that the function will panic if a wrong type series is given
	assert.PanicsWithError(t, "only series with float type can be rounded", func() {
		NewSeries("col", "foo", "bar").Round(2, true)
	})

	// assert that the function will panic if there's an invalid value
	s.Data[2] = "foo"
	_, err := strconv.ParseFloat("foo", 64)
	assert.PanicsWithError(t, fmt.Sprintf("Invalid value at row %d. %s", 2, err.Error()), func() {
		s.Round(2, true)
	})
}
