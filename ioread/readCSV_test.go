package ioread

import (
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	actual, _ := ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisSample.csv")})
	assert.Equal(t, expected, actual)

	// assert that index col is set correctly
	expected = expected.ShallowCopy().SetIndex("petal_length")
	actual, _ = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisSample.csv"), IndexCol: "petal_length"})
	assert.Equal(t, expected, actual)

	// assert that an error will be produced for erroneous data
	actual, err = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisIncorrect.csv")})
	assert.Nil(t, actual)
	assert.EqualError(t, err, "record on line 4: wrong number of fields")

	// assert that error will be returned for malformed dates
	actual, err = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "dateIncorrectSample.csv"),
		DateCols: []string{"date"}, DateFormat: "2006-01-02"})
	assert.Nil(t, actual)
	assert.NotNil(t, err)

	// assert that function skip error lines correctly
	actual, err = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisIncorrect.csv"),
		SkipErrorLines: true})
	assert.NotNil(t, actual)
	assert.Nil(t, err)

	// assert that function warns about error lines correctly
	rescueStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	_, _ = ReadCSV(CsvOptions{Path: filepath.Join(dataPath, "irisIncorrect.csv"), SkipErrorLines: true,
		WarnErrorLines: true})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdOut
	assert.Equal(t, "record on line 4: wrong number of fields\n", string(out))
}
