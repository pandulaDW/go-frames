package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_IsNa(t *testing.T) {
	s := NewSeries("col", 12, 43, nil, 1.2, nil, 90)

	// assert that na values are correctly identified
	assert.Equal(t, s.IsNa(), NewSeries("isna(col)", false, false, true, false, true, false))
}

func TestSeries_NotNa(t *testing.T) {
	s := NewSeries("col", 12, 43, nil, 1.2, nil, 90)

	// assert that na values are correctly identified
	assert.Equal(t, s.NotNa(), NewSeries("notna(col)", true, true, false, true, false, true))
}

func TestSeries_CountOfNA(t *testing.T) {
	s := NewSeries("col", 12, 43, nil, 1.2, "", 90)

	// assert that na values are correctly counted
	assert.Equal(t, s.CountOfNA(), 2)
}

func TestSeries_IsBlank(t *testing.T) {
	// assert that function returns an error if invalid series is entered
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		s.IsBlank()
	})

	// assert that na values are correctly identified
	s = NewSeries("col", "foo", "bar", "", "baz", "", "bar")
	assert.Equal(t, s.IsBlank(), NewSeries("isBlank(col)", false, false, true, false, true, false))
}

func TestSeries_NotBlank(t *testing.T) {
	// assert that function returns an error if invalid series is entered
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		s.NotBlank()
	})

	// assert that na values are correctly identified
	s = NewSeries("col", "foo", "bar", "", "baz", "", "bar")
	assert.Equal(t, s.NotBlank(), NewSeries("notBlank(col)", true, true, false, true, false, true))
}
