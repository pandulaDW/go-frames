package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSeries_RegexExtract(t *testing.T) {
	re := regexp.MustCompile("foo.+")

	// assert that function panics at incorrect DType
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		s.RegexContains(re)
	})

	s = NewSeries("test", "foo", "food", "foodies", nil, "food and dining")

	// assert that function panics if invalid values are given
	s.Data[2] = 12
	assert.PanicsWithError(t, errors.InvalidSeriesValError(12, 2, "test").Error(), func() {
		s.RegexContains(re)
	})

	// assert that function extract values correctly
	s.Data[2] = "foodies"
	expected := NewSeries("regex_extract(test)", false, true, true, false, true)
	assert.Equal(t, expected, s.RegexContains(re))
}
