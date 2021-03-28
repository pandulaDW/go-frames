package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSeries_RegexContains(t *testing.T) {
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
	expected := NewSeries("regex_contains(test)", false, true, true, false, true)
	assert.Equal(t, expected, s.RegexContains(re))
}

func TestSeries_RegexExtract(t *testing.T) {
	re := regexp.MustCompile(`(\w+/\w+)`)

	// assert that function panics at incorrect DType
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		s.RegexExtract(re, 1)
	})

	s = NewSeries("test", "swede foo/bar sd", "swede", "fill/bill kill", nil, "matter na/cl size")

	// assert that function panics if invalid values are given
	s.Data[2] = 12
	assert.PanicsWithError(t, errors.InvalidSeriesValError(12, 2, "test").Error(), func() {
		s.RegexExtract(re, 1)
	})

	// assert that function panics if index is out of range values correctly
	s.Data[2] = "fill/bill kill"
	assert.PanicsWithError(t, errors.CustomError("index is out of range").Error(), func() {
		s.RegexExtract(re, 3)
	})

	// assert that the function extract values correctly
	expected := NewSeries("regex_extract(test)", "foo/bar", "", "fill/bill", "", "na/cl")
	assert.Equal(t, expected, s.RegexExtract(re, 1))
}

func TestSeries_RegexReplace(t *testing.T) {
	re := regexp.MustCompile(`\d`)

	// assert that function panics at incorrect DType
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		s.RegexReplace(re, "|")
	})

	s = NewSeries("test", "foo 2 bar", "bar", nil, "23 jump street", "foo 2 bar 2")

	// assert that function panics if invalid values are given
	s.Data[2] = 12
	assert.PanicsWithError(t, errors.InvalidSeriesValError(12, 2, "test").Error(), func() {
		s.RegexReplace(re, "|")
	})

	// assert that the function replaces values correctly
	s.Data[2] = nil
	expected := NewSeries("regex_replace(test)", "foo | bar", "bar", "", "|| jump street", "foo | bar |")
	assert.Equal(t, expected, s.RegexReplace(re, "|"))
}
