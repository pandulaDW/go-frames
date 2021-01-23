package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDataFrame(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	col5 := series.NewSeries("col5", 14, 12.23, 32.5, 64, 34.1)

	expected := NewDataFrame(col1, col2, col3, col4, col5)
	actual := &DataFrame{
		Data: map[string]*series.Series{"col1": col1, "col2": col2, "col3": col3,
			"col4": col4, "col5": col5},
		length: 5,
		columns: []*base.Column{{Name: "col1", Dtype: base.Int, ColIndex: 0},
			{Name: "col2", Dtype: base.Object, ColIndex: 1},
			{Name: "col3", Dtype: base.Float, ColIndex: 2},
			{Name: "col4", Dtype: base.Bool, ColIndex: 3},
			{Name: "col5", Dtype: base.Float, ColIndex: 4}},
	}

	// assert that the dataframe is created successfully
	assert.Equal(t, expected, actual, "dataframe is created successfully")

	// assert that an empty dataframe is returned when no data is provided
	assert.Equal(t, new(DataFrame), NewDataFrame(),
		"empty dataframe is returned when calling with no data")

	// assert that the function panics when mismatched row numbers are given
	assert.PanicsWithError(t, "mismatched row lengths found. Dataframe can only contain equal number of rows",
		func() {
			col6 := series.NewSeries("col1", 12, 34, 54)
			NewDataFrame(col1, col2, col3, col4, col5, col6)
		})
}
