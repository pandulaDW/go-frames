package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
)

// ConvertRowContentToDF will convert the row based content to a dataframe
func ConvertRowContentToDF(colNames []string, content [][]string) *DataFrame {
	seriesSlice := make([]*series.Series, 0)

	for colIndex, colName := range colNames {
		colData := make([]interface{}, 0)
		for rowIndex := 0; rowIndex < len(content); rowIndex++ {
			colData = append(colData, content[rowIndex][colIndex])
		}

		s := series.NewSeries(colName, colData...)
		seriesSlice = append(seriesSlice, s)
	}

	return NewDataFrame(seriesSlice...)
}
