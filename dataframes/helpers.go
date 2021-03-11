package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
)

// ConvertRowContentToDF will convert the row based content to a dataframe
func ConvertRowContentToDF(colNames []string, content [][]string) *DataFrame {
	seriesSlice := make([]*series.Series, 0, len(colNames))
	ch := make(chan *series.Series, len(colNames))

	for colIndex, colName := range colNames {
		colData := make([]interface{}, 0)
		for rowIndex := 0; rowIndex < len(content); rowIndex++ {
			colData = append(colData, content[rowIndex][colIndex])
		}

		go func(colName string, colData []interface{}) {
			s := series.NewSeries(colName, colData...)
			ch <- s
		}(colName, colData)
	}

	count := 0
	for s := range ch {
		seriesSlice = append(seriesSlice, s)
		count++
		if count == cap(ch) {
			close(ch)
		}
	}

	return NewDataFrame(seriesSlice...)
}
