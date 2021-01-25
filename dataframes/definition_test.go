package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFrame_Length(t *testing.T) {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	df := NewDataFrame(col1, col2)

	// assert that the function returns the correct dataframe length
	assert.Equal(t, 5, df.Length())
}
