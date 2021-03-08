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

func TestDateParsing(t *testing.T) {
	df := dataframes.NewDataFrame(series.NewSeries("col", 12, 45, 6))

	// assert that function returns nil if arguments are not present
	options := &CsvOptions{}
	assert.Nil(t, dateParsing(options, df))
}
