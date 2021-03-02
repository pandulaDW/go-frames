package ioread

import (
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestReadCSV(t *testing.T) {
	wd, _ := os.Getwd()
	dataPath := filepath.Join(wd, "..", "data")
	expected := dataframes.NewDataFrame(
		series.NewSeries("sepal_length", 5.1, 4.9, 4.7, 4.6),
		series.NewSeries("sepal_width", 3.5, 3.0, 3.2, 3.1),
		series.NewSeries("petal_length", 1.4, 1.4, 1.3, 1.5),
		series.NewSeries("petal_width", 0.2, 0.2, 0.2, 0.2),
		series.NewSeries("species", "setosa", "setosa", "setosa", "setosa"))

	// assert that an error is returned if there's an error in reading the file
	file, err := ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "testFile.csv")})
	assert.NotNil(t, err)
	assert.Nil(t, file)

	// assert that a dataframe is created correctly
	actual, _ := ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisSample.csv"), Delimiter: ","})
	assert.Equal(t, expected, actual)

	// assert that index col is set correctly
	expected = expected.ShallowCopy().SetIndex("petal_length")
	actual, _ = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisSample.csv"), Delimiter: ",",
		IndexCol: "petal_length"})
	assert.Equal(t, expected, actual)

	// assert that an error will be produced for mismatched samples
	actual, err = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisIncorrect.csv"), Delimiter: ","})
	assert.Nil(t, actual)
	assert.EqualError(t, err, "mismatched number of columns in row 3")
}
