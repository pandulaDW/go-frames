package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
)

// seriesHolder will hold series object pointers with their indexes
type seriesHolder struct {
	s     *series.Series
	index int
}

// ConvertRowContentToDF will convert the row based content to a dataframe
func ConvertRowContentToDF(colNames []string, content [][]string) *DataFrame {
	ch := make(chan seriesHolder, len(colNames))

	for colIndex, colName := range colNames {
		colData := make([]interface{}, 0)
		for rowIndex := 0; rowIndex < len(content); rowIndex++ {
			colData = append(colData, content[rowIndex][colIndex])
		}

		go func(colName string, colData []interface{}, index int) {
			sh := seriesHolder{s: series.NewSeries(colName, colData...), index: index}
			ch <- sh
		}(colName, colData, colIndex)
	}

	seriesSlice := make([]seriesHolder, 0, len(colNames))
	count := 0
	for sh := range ch {
		seriesSlice = append(seriesSlice, sh)
		count++
		if count == cap(ch) {
			close(ch)
		}
	}

	// create the ordered series slice
	orderedSeriesSlice := make([]*series.Series, len(colNames))
	for _, sh := range seriesSlice {
		orderedSeriesSlice[sh.index] = sh.s
	}

	return NewDataFrame(orderedSeriesSlice...)
}

// ConvertMapToDataFrame will convert a Go map to a two column DataFrame by using the keys in one column and
// the corresponding values in the other column. The map can have any type of a key and value.
//
// The two columns will be named as keys and values.
func ConvertMapToDataFrame(m map[interface{}]interface{}) *DataFrame {
	keyData := make([]interface{}, 0, len(m))
	valueData := make([]interface{}, 0, len(m))

	for key, val := range m {
		keyData = append(keyData, key)
		valueData = append(valueData, val)
	}

	keyS := series.NewSeries("keys", keyData...)
	valueS := series.NewSeries("values", valueData...)

	return NewDataFrame(keyS, valueS)
}
