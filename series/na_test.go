package series

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_IsNa(t *testing.T) {
	s := NewSeries("col", 12, 43, "", 1.2, "", 90)

	// assert that na values are correctly identified
	assert.Equal(t, s.IsNa(), NewSeries("isna(col)", false, false, true, false, true, false))
}

func TestSeries_NotNa(t *testing.T) {
	s := NewSeries("col", 12, 43, "", 1.2, "", 90)

	// assert that na values are correctly identified
	assert.Equal(t, s.NotNa(), NewSeries("notna(col)", true, true, false, true, false, true))
}
