package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFrame_Agg(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", 54.31, 1.23, 45.6, 23.12, 23.2)
	col3 := series.NewSeries("col3", 14, 124.23, 32, 64.65, 34)
	df := NewDataFrame(col1, col2, col3)

	// assert that the function panics if an unknown column name is given
	assert.PanicsWithError(t, "test not found in the dataframe", func() {
		df.Agg([]string{"col1", "test"}, base.AVG)
	})

	// assert that the function panics if an unknown aggregator is given
	assert.PanicsWithError(t, "test is not a valid aggregator identifier", func() {
		df.Agg([]string{"col1", "col2"}, "test")
	})

	// assert that the max aggregator works as expected
	assert.Equal(t, []interface{}{90, 54.31, 124.23}, df.Agg([]string{"col1", "col2", "col3"}, base.MAX))

	// assert that the min aggregator works as expected
	assert.Equal(t, []interface{}{12, 1.23, float64(14)}, df.Agg([]string{"col1", "col2", "col3"}, base.MIN))

	// assert that the sum aggregator works as expected
	assert.Equal(t, []interface{}{float64(255), 147.46, 268.88}, df.Agg([]string{"col1", "col2", "col3"}, base.SUM))

	// assert that the avg aggregator works as expected
	assert.Equal(t, []interface{}{float64(255) / 5, 147.46 / 5, 268.88 / float64(5)},
		df.Agg([]string{"col1", "col2", "col3"}, base.AVG))
}
