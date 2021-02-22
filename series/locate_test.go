package series

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_Loc(t *testing.T) {
	s := NewSeries("col", 89, 69.1, 2.34, 1.58, 2.4, 23, 54, 10)

	// assert that the function panics when out of range index is given
	assert.PanicsWithError(t, "index 9 is out of range", func() {
		s.Loc([]int{1, 5, 2, 9})
	})

	// assert that the function panics when a negative index is given
	assert.PanicsWithError(t, "index -2 is out of range", func() {
		s.Loc([]int{1, 5, -2, 4})
	})

	// assert that the function returns a series correctly
	assert.Equal(t, NewSeries("col", 2.34, 2.4, 10), s.Loc([]int{2, 4, 7}))
}
