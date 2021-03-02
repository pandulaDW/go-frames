package ioread

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileHandling(t *testing.T) {
	// assert that function panics with error when file cannot be read properly
	_, err := fileHandling("data/testFile")
	assert.EqualError(t, err,
		"error in reading the file: \nopen data/testFile: The system cannot find the path specified.")

	// assert that function returns a file handler for a correct file
	file, _ := fileHandling("commonFuncs.go")
	assert.Equal(t, "*os.File", fmt.Sprintf("%T", file))
}

func TestConvertRowContentToDF(t *testing.T) {
	content := make([][]string, 3)
	content[0] = []string{"1.2", "foo", "True"}
	content[1] = []string{"3.4", "bar", "True"}
	content[2] = []string{"4.5", "baz", "False"}
	colNames := []string{"floatCol", "strCol", "boolCol"}

	expected := dataframes.NewDataFrame(
		series.NewSeries(colNames[0], "1.2", "3.4", "4.5"),
		series.NewSeries(colNames[1], "foo", "bar", "baz"),
		series.NewSeries(colNames[2], "True", "True", "False"))

	// assert that row content is converted to a dataframe successfully
	assert.Equal(t, expected, convertRowContentToDF(colNames, content))
}
