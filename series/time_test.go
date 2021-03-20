package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelperTimeMethods(t *testing.T) {
	// assert that function returns an error if invalid series is entered
	s := NewSeries("test", 12, 43, 11, 10)
	assert.PanicsWithError(t, errors.IncorrectDataType(base.DateTime).Error(), func() {
		helperTimeMethods(s, "YEAR")
	})

	s = NewSeries("test", "2016-01-02", "2016-01-03", "2016-01-04")
	_ = s.CastAsTime("2006-01-02")
	s.Data[2] = 5
	assert.PanicsWithError(t, errors.InvalidSeriesValError(5, 2, s.column.Name).Error(), func() {
		helperTimeMethods(s, "YEAR")
	})
}

func TestSeries_Year(t *testing.T) {
	// assert that year is extracted properly
	s := NewSeries("test", "2016-01-02", "2017-01-03", "2021-01-04")
	_ = s.CastAsTime("2006-01-02")
	assert.Equal(t, NewSeries("year(test)", 2016, 2017, 2021), s.Year())
}

func TestSeries_Month(t *testing.T) {
	// assert that month is extracted properly
	s := NewSeries("test", "2016-03-02", "2017-05-03", "2021-11-04")
	_ = s.CastAsTime("2006-01-02")
	assert.Equal(t, NewSeries("month(test)", "March", "May", "November"), s.Month())
}

func TestSeries_Day(t *testing.T) {
	// assert that year is extracted properly
	s := NewSeries("test", "2016-01-02", "2017-01-03", "2021-01-24")
	_ = s.CastAsTime("2006-01-02")
	assert.Equal(t, NewSeries("day(test)", 2, 3, 24), s.Day())
}
