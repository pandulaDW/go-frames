package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFrame_FilterBySeries(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", 54.31, 1.23, 45.6, 23.12, 23.2)
	df := NewDataFrame(col1, col2)

	// assert that function panics if wrong data type is given
	assert.PanicsWithError(t, errors.IncorrectDataType(base.Bool).Error(), func() {
		df.FilterBySeries(col1)
	})

	// assert that function panics if mismatched number of elements in the series
	assert.PanicsWithError(t, errors.MismatchedNumOfRows(df.length, 4).Error(), func() {
		df.FilterBySeries(series.NewSeries("col", true, false, false, true))
	})

	// assert that function panics if malformed bool value is given
	s := series.NewSeries("col", true, false, true, false, false)
	s.Data[2] = "foo"
	assert.PanicsWithError(t, errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		df.FilterBySeries(s)
	})

	// assert that function returns correctly filtered data
	expected := NewDataFrame(series.NewSeries("col1", 12, 54, 90),
		series.NewSeries("col2", 54.31, 45.6, 23.2))
	assert.Equal(t, expected, df.FilterBySeries(series.NewSeries("col", true, false, true, false, true)))
}
