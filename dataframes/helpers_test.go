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

func TestConvertMapToDataFrame(t *testing.T) {
	m := make(map[interface{}]interface{})
	m["foo"] = 12
	m["bar"] = 15
	m["baz"] = 16

	// assert that the function returns the correct dataframe
	expected := NewDataFrame(series.NewSeries("keys", "foo", "bar", "baz"),
		series.NewSeries("values", 12, 15, 16))
	assert.Equal(t, expected, ConvertMapToDataFrame(m))
}
