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

func TestSeries_Lower(t *testing.T) {
	// assert that function returns correctly lowered series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("test", "foo", "bar", "baz"), s.Lower())
}

func TestSeries_Upper(t *testing.T) {
	// assert that function returns correctly upper series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("test", "FOO", "BAR", "BAZ"), s.Upper())
}

func TestSeries_Capitalized(t *testing.T) {
	// assert that function returns correctly capitalized series
	s := NewSeries("test", "foo", "BAR", "BaZ")
	assert.Equal(t, NewSeries("test", "Foo", "BAR", "BaZ"), s.Capitalized())
}

func TestSeries_Trim(t *testing.T) {
	// assert that function returns correctly trimmed series
	s := NewSeries("test", "foo ", " BAR", "  BaZ  ", "march")
	assert.Equal(t, NewSeries("test", "foo", "BAR", "BaZ", "march"), s.Trim())
}
