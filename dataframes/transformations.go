package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/series"
)

// Transpose creates a new dataframe with the current dataframe being transposed.
// If header is set to true, columns of the current dataframe will be preserved as the
// first column of the transposed column.
//
// Column names will be added incrementally and these can be changed using df.SetColumnName() method
func (df *DataFrame) Transpose(header bool) *DataFrame {
	newSeriesArray := make([]*series.Series, 0)

	if header {
		colNames := make([]interface{}, 0)
		for _, col := range df.columns {
			colNames = append(colNames, col.Name)
		}
		newSeriesArray = append(newSeriesArray, series.NewSeries("v0", colNames...))
	}

	for i := 0; i < df.length; i++ {
		colData := make([]interface{}, 0)
		for _, col := range df.columns {
			colData = append(colData, df.Data[col.Name].Data[i])
		}

		newSeries := series.NewSeries(fmt.Sprintf("v%d", i+1), colData...)
		newSeriesArray = append(newSeriesArray, newSeries)
	}

	return NewDataFrame(newSeriesArray...)
}
