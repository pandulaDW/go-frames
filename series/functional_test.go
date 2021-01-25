package series

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSeries_Map(t *testing.T) {
	mapper := func(a interface{}) (interface{}, error) {
		floatVal, ok := a.(float64)
		if !ok {
			return nil, errors.New("only float values can be included")
		}
		return math.Pow(floatVal, 2), nil
	}

	s1 := NewSeries("col", 12.23, 23.11, 14.65, 6.33, 11.90)
	s2 := NewSeries("col", 12, 23, 14, 6, 11)
	expected := NewSeries("col", 149.5729, 534.0721, 214.6225, 40.0689, 141.61)

	// assert that the squared function mapper works for an float series
	actual, err := s1.Map(mapper)

	// assert that mapper returns correct values
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)

	actual, err = s2.Map(mapper)
	// assert that the squared function mapper doesn't work for a float series
	assert.Nil(t, actual)
	assert.EqualError(t, err, "only float values can be included")
}
