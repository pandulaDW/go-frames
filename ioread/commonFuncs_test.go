package ioread

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/errors"
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
	df := dataframes.NewDataFrame(
		series.NewSeries("col1", 12, 45, 6),
		series.NewSeries("col2", "2013-11-05", "2021-10-15", "2015-12-11"))
	format := "2006-01-02"

	// assert that function returns nil if arguments are not present
	options := &CsvOptions{}
	assert.Nil(t, dateParsing(options, df))

	// assert that function returns an error if cols are given and format is not given
	options = &CsvOptions{DateCols: []string{"col"}}
	assert.EqualError(t,
		dateParsing(options, df), "DateFormat field should not be empty if DateCols field is present")

	// assert that function returns an error if column is not included
	options = &CsvOptions{DateCols: []string{"col3"}, DateFormat: format}
	assert.EqualError(t, dateParsing(options, df), errors.ColumnNotFound("col3").Error())

	// assert that function returns a cast error for a parsing issue
	options = &CsvOptions{DateCols: []string{"col1", "col2"}, DateFormat: format}
	assert.EqualError(t, dateParsing(options, df),
		"only a series with object type can be inferred as a datetime series")

	// assert that function returns nil if no errors are found
	options = &CsvOptions{DateCols: []string{"col2"}, DateFormat: format}
	assert.Nil(t, dateParsing(options, df))
}
