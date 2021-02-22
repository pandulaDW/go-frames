package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
)

//Loc will access a group of rows by an integer array and and columns by an string array.
// Panics if out of range indices are found or an undefined column is given.
//
// To return all the columns, use df.Loc([row_indices], df.Columns())
func (df *DataFrame) Loc(indices []int, columns []string) *DataFrame {
	seriesArray := make([]*series.Series, 0)

	for _, col := range columns {
		s, ok := df.Data[col]
		if !ok {
			panic(errors.CustomError(col + " column is not found"))
		}
		dataSlice := s.Loc(indices) // will panic here if indices are out of range
		newSeries := series.NewSeries(col, dataSlice...)
		seriesArray = append(seriesArray, newSeries)
	}

	newDF := NewDataFrame(seriesArray...)
	return newDF
}
