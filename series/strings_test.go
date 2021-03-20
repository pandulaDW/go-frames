package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelperStringMethods(t *testing.T) {
	// assert that function returns an error if invalid series is entered
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		helperStringMethods(s, func(val string) string {
			return ""
		})
	})

	s = NewSeries("test", "foo", "bar", "baz")
	s.Data[2] = 5
	assert.PanicsWithError(t, errors.InvalidSeriesValError(5, 2, s.column.Name).Error(), func() {
		helperStringMethods(s, func(val string) string {
			return ""
		})
	})
}

func TestHelperStringBooleanMethods(t *testing.T) {
	// assert that function returns an error if invalid series is entered
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Object).Error(), func() {
		helperStringBooleanMethods(s, "test", func(val1, val2 string) bool {
			return true
		})
	})

	s = NewSeries("test", "foo", "bar", "baz")
	s.Data[2] = 5
	assert.PanicsWithError(t, errors.InvalidSeriesValError(5, 2, s.column.Name).Error(), func() {
		helperStringBooleanMethods(s, "test", func(val1, val2 string) bool {
			return true
		})
	})
}

func TestSeries_Lower(t *testing.T) {
	// assert that function returns correctly lowered series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("ToLower(test)", "foo", "bar", "baz"), s.Lower())
}

func TestSeries_Upper(t *testing.T) {
	// assert that function returns correctly upper series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("ToUpper(test)", "FOO", "BAR", "BAZ"), s.Upper())
}

func TestSeries_Capitalized(t *testing.T) {
	// assert that function returns correctly capitalized series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("Title(test)", "Foo", "BAR", "BaZ"), s.Capitalized())
}

func TestSeries_Trim(t *testing.T) {
	// assert that function returns correctly trimmed series
	s := NewSeries("test", "foo ", " BAR", "  BaZ  ", "march")
	assert.Equal(t, NewSeries("TrimSpace(test)", "foo", "BAR", "BaZ", "march"), s.Trim())
}

func TestSeries_Contains(t *testing.T) {
	// assert that function returns a correct bool Series
	s := NewSeries("test", "we are", "leaving", "right now", "now now!!")
	assert.Equal(t, NewSeries("Contains(test)", false, false, true, true), s.Contains("now"))
}

func TestSeries_StartsWith(t *testing.T) {
	// assert that function returns a correct bool Series
	s := NewSeries("test", "foo", "bar is", "food now", "now now!!")
	assert.Equal(t, NewSeries("HasPrefix(test)", true, false, true, false), s.StartsWith("foo"))
}

func TestSeries_EndsWith(t *testing.T) {
	// assert that function returns a correct bool Series
	s := NewSeries("test", "we are", "leaving now", "right now", "bar")
	assert.Equal(t, NewSeries("HasSuffix(test)", false, true, true, false), s.EndsWith("now"))
}
