package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFrame_Transpose(t *testing.T) {
	df := NewDataFrame(
		series.NewSeries("col1", 12, 34, 54),
		series.NewSeries("col2", "foo", "bar", "raz"),
		series.NewSeries("col3", 54.31, 1.23, 45.6))

	tCol1 := series.NewSeries("v1", 12, "foo", 54.31)
	tCol2 := series.NewSeries("v2", 34, "bar", 1.23)
	tCol3 := series.NewSeries("v3", 54, "raz", 45.6)

	expected := NewDataFrame(tCol1, tCol2, tCol3)

	// assert that the transpose method works as expected without header column
	assert.Equal(t, expected, df.Transpose(false))

	// assert that the transpose method works as expected with header column
	expected = NewDataFrame(series.NewSeries("v0", "col1", "col2", "col3"), tCol1, tCol2, tCol3)
	assert.Equal(t, expected, df.Transpose(true))
}
