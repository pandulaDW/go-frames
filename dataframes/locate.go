package dataframes

import "github.com/pandulaDW/go-frames/series"

//Loc will access a group of rows by an integer array and and columns by an string array.
// Panics if out of range indices are found or a undefined column is given.
//
// To return all columns, use df.Loc([row_indices], df.Columns())
func (df *DataFrame) Loc(indices []int, columns []string) *DataFrame {
	seriesArray := make([]*series.Series, 0)
	_ = seriesArray

	//for _, index := range indices {
	//
	//}

	return nil
}
