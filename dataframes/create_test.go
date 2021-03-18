package dataframes

import (
	"fmt"
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
	col6 := series.NewSeries("col6", 10, 12, "", 45, 89)

	expected := NewDataFrame(col1, col2, col3, col4, col5, col6)
	actual := &DataFrame{
		data: map[string]*series.Series{"col1": col1, "col2": col2, "col3": col3,
			"col4": col4, "col5": col5, "col6": col6},
		length: 5,
		columns: []*base.Column{{Name: "col1", Dtype: base.Int, ColIndex: 0},
			{Name: "col2", Dtype: base.Object, ColIndex: 1},
			{Name: "col3", Dtype: base.Float, ColIndex: 2},
			{Name: "col4", Dtype: base.Bool, ColIndex: 3},
			{Name: "col5", Dtype: base.Float, ColIndex: 4},
			{Name: "col6", Dtype: base.Int, ColIndex: 5}},
		Index: Index{
			Data:     series.NewSeries("#", 0, 1, 2, 3, 4),
			IsCustom: false,
		},
	}

	for _, column := range actual.columns {
		actual.data[column.Name].GetColumn().ColIndex = column.ColIndex
	}

	// assert that the dataframe is created successfully
	assert.Equal(t, expected, actual)

	// assert that an empty dataframe is returned when no data is provided
	assert.Equal(t, new(DataFrame), NewDataFrame())

	// assert that the function panics when mismatched row numbers are given
	assert.PanicsWithError(t, "mismatched number of rows provided. requires 5 rows, but 3 was provided",
		func() {
			newCol := series.NewSeries("col1", 12, 34, 54)
			NewDataFrame(col1, col2, col3, col4, col5, newCol)
		})

	// assert that the function converts int values to floats in mixed type series
	assert.Equal(t, []interface{}{float64(14), 12.23, 32.5, float64(64), 34.1}, col5.Data)

	// assert that the function panics when duplicate column is given
	assert.PanicsWithError(t, "col1 column is already in the dataframe",
		func() {
			newCol := series.NewSeries("col1", 12, 34, 54, 67, 89)
			NewDataFrame(col1, col2, col3, col4, col5, newCol)
		})
}

func TestDataFrame_DeepCopy(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	df := NewDataFrame(col1, col2, col3)
	copied := df.DeepCopy()

	// assert that two object references are different
	assert.NotEqual(t, fmt.Sprintf("%p", df), fmt.Sprintf("%p", copied))

	// assert that the dataframe objects are equal
	assert.Equal(t, df, copied)

	// assert that series object references are not equal
	assert.NotEqual(t, fmt.Sprintf("%p", col1), fmt.Sprintf("%p", copied.data["col1"]))
	assert.NotEqual(t, fmt.Sprintf("%p", col2), fmt.Sprintf("%p", copied.data["col2"]))
	assert.NotEqual(t, fmt.Sprintf("%p", col3), fmt.Sprintf("%p", copied.data["col3"]))
}

func TestDataFrame_ShallowCopy(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	df := NewDataFrame(col1, col2, col3)
	copied := df.ShallowCopy()

	// assert that two object references are different
	assert.NotEqual(t, fmt.Sprintf("%p", df), fmt.Sprintf("%p", copied))

	// assert that the dataframe objects are equal
	assert.Equal(t, df, copied)
}

func TestDataFrame_IsEqual(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")

	df := NewDataFrame(col1, col2)
	df2 := df

	// assert that function returns true if, both dataframes pointers are same
	assert.Equal(t, true, df.IsEqual(df2))

	// assert that function returns false if, both dataframes pointers are not the same
	assert.Equal(t, false, df.IsEqual(NewDataFrame(col1, col2)))
}

func TestDataFrame_IsDeepEqual(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")

	df1 := NewDataFrame(col1, col2)
	df2 := NewDataFrame(col1, col2)

	// assert that function returns true if both dataframes are same
	assert.Equal(t, true, df1.IsDeepEqual(df2))

	// assert that function returns false if both dataframes are not equal
	col2 = series.NewSeries("col2", "foo", "bar", "raz", "apple", "mangoes")
	df3 := NewDataFrame(col1, col2)
	assert.Equal(t, false, df1.IsDeepEqual(df3))
}
