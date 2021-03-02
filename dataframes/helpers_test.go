package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertRowContentToDF(t *testing.T) {
	content := make([][]string, 3)
	content[0] = []string{"1.2", "foo", "True"}
	content[1] = []string{"3.4", "bar", "True"}
	content[2] = []string{"4.5", "baz", "False"}
	colNames := []string{"floatCol", "strCol", "boolCol"}

	expected := NewDataFrame(
		series.NewSeries(colNames[0], "1.2", "3.4", "4.5"),
		series.NewSeries(colNames[1], "foo", "bar", "baz"),
		series.NewSeries(colNames[2], "True", "True", "False"))

	// assert that row content is converted to a dataframe successfully
	assert.Equal(t, expected, ConvertRowContentToDF(colNames, content))
}
