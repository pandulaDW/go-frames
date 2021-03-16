package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
)

// AddColumn adds a new column to the existing dataframe. To create a new dataframe object,
// create a shallow copy first before adding the series.
//
// Panics if the length of the series is not equal to the length of the dataframe.
func (df *DataFrame) AddColumn(s *series.Series) *DataFrame {
	if s.Len() != df.Length() {
		panic(errors.MismatchedNumOfRows(df.length, s.Len()))
	}

	// append column
	df.columns = append(df.columns, s.GetColumn())

	// add the data
	df.data[s.GetColumn().Name] = s

	return df
}
