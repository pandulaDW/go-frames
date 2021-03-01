package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
)

// ILoc will access a group of rows by an integer array and and columns by an string array.
// Panics if out of range indices are found or an undefined column is given.
//
// To return all the columns, use df.ILoc([row_indices], df.Columns())
func (df *DataFrame) ILoc(indices []int, columns []string) *DataFrame {
	seriesArray := make([]*series.Series, 0)

	for _, col := range columns {
		s, ok := df.Data[col]
		if !ok {
			panic(errors.CustomError(col + " column is not found"))
		}
		newSeries := s.Loc(indices) // will panic here if indices are out of range
		seriesArray = append(seriesArray, newSeries)
	}

	newDF := NewDataFrame(seriesArray...)
	return newDF
}

// Head function returns the first n rows for the dataframe. It is useful for quickly
// testing if your dataframe has the right type of data in it.
//
// Panics if n is higher than the length of the dataframe
func (df *DataFrame) Head(n int) *DataFrame {
	if n > df.length {
		panic(errors.CustomError("n cannot be higher than the length of the dataframe"))
	}

	return df.ILoc(helpers.Range(0, n, 1), df.Columns())
}

// Tail function returns last n rows from the object based on position. It is useful
// for quickly verifying data, for example, after sorting or appending rows.
//
// Panics if n is higher than the length of the dataframe
func (df *DataFrame) Tail(n int) *DataFrame {
	if n > df.Length() {
		panic(errors.CustomError("n cannot be higher than the length of the dataframe"))
	}

	indices := helpers.Range(df.length-1, df.length-n-1, -1)
	reversedIndices := helpers.ReverseArray(helpers.ToInterfaceFromInt(indices))
	return df.ILoc(helpers.ToIntArray(reversedIndices), df.Columns())
}
