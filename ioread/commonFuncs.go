package ioread

import (
	"github.com/pandulaDW/go-frames/dataframes"
	errors2 "github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
	"os"
)

//fileHandling will handle the file and will return the file
func fileHandling(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(errors2.CustomWithStandardError("error in reading the file", err))
	}

	return file
}

//convertRowContentToDF will convert the row based content to a dataframe
func convertRowContentToDF(colNames []string, content [][]string) *dataframes.DataFrame {
	seriesSlice := make([]*series.Series, 0)

	for colIndex, colName := range colNames {
		colData := make([]interface{}, 0)
		for rowIndex := 0; rowIndex < len(content); rowIndex++ {
			colData = append(colData, content[rowIndex][colIndex])
		}

		s := series.NewSeries(colName, colData...)
		seriesSlice = append(seriesSlice, s)
	}

	return dataframes.NewDataFrame(seriesSlice...)
}
