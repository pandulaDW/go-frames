package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
)

// SetIndex sets the given column as the index. Panics if the column name is not found
func (df *DataFrame) SetIndex(colName string) *DataFrame {
	if helpers.LinearSearch(colName, helpers.ToInterfaceFromString(df.Columns())) == -1 {
		panic(errors.CustomError("column not found"))
	}

	// set the column as the index
	df.Index = Index{Data: df.Data[colName], IsCustom: true}

	// delete series from the column map
	delete(df.Data, colName)

	// filter out from the columns array
	filteredCols := make([]*base.Column, 0)
	for _, column := range df.columns {
		if column.Name != colName {
			filteredCols = append(filteredCols, column)
		}
	}

	// set the filtered
	df.columns = filteredCols

	return df
}

// SetIndexWithSeries sets the index as the given series
func (df *DataFrame) SetIndexWithSeries(s *series.Series) *DataFrame {
	return df
}
