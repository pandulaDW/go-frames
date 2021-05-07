package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSeries_Apply(t *testing.T) {
	mapper := func(a interface{}) (interface{}, error) {
		floatVal, ok := a.(float64)
		if !ok {
			return nil, fmt.Errorf("%s", "only float values can be included")
		}
		return math.Pow(floatVal, 2), nil
	}

	s1 := NewSeries("col", 12.23, 23.11, 14.65, nil, 6.33, 11.90)
	s2 := NewSeries("col", 12, 23, 14, 6, 11)
	expected := NewSeries("func1(col)", 149.5729, 534.0721, 214.6225, nil, 40.0689, 141.61)

	// assert that the squared function mapper works for an float series
	actual, err := s1.Apply(mapper)

	// assert that mapper returns correct values
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)

	actual, err = s2.Apply(mapper)
	// assert that the squared function mapper doesn't work for a float series
	assert.Nil(t, actual)
	assert.EqualError(t, err, "only float values can be included")
}

func TestCompose(t *testing.T) {
	s1 := NewSeries("col1", true, false, true, true, false)
	s2 := NewSeries("col2", false, true, true, false, false)

	// assert that function panics with mismatched len series will
	assert.PanicsWithError(t, `col3 is invalid. mismatched number of rows provided. requires 5 rows, but 3 was provided`,
		func() {
			compose("AND", "test", s1, s2, NewSeries("col3", false, true, true))
		})

	// assert that function panics if invalid series val is given
	s3 := NewSeries("col3", false, 12, true, false, false)
	assert.PanicsWithError(t, errors.InvalidSeriesValError(s3, 1, "col3").Error(), func() {
		compose("AND", "test", s1, s2, s3)
	})
}

func TestComposeWithAnd(t *testing.T) {
	s1 := NewSeries("col1", true, false, true, true, false)
	s2 := NewSeries("col2", false, true, true, false, false)
	expected := NewSeries("compose-with-and", false, false, true, false, false)

	// assert that function correctly compose bool values
	assert.Equal(t, expected, ComposeWithAnd(s1, s2))
}

func TestComposeWithOR(t *testing.T) {
	s1 := NewSeries("col1", true, false, true, false, false)
	s2 := NewSeries("col2", false, true, true, false, false)
	s3 := NewSeries("col2", true, true, false, false, true)
	expected := NewSeries("compose-with-or", true, true, true, false, true)

	// assert that function correctly compose bool values
	assert.Equal(t, expected, ComposeWithOR(s1, s2, s3))
}
