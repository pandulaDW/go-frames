package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
)

// AddColumn adds a new column to the existing dataframe. The first argument is the
// series to be added and the inplace argument specifies whether a new dataframe is returned
// or the new column will be added to the existing dataframe. Either way, the modified or the new
// dataframe will be returned from the function
//
// Panics if the length of the series is not equal to the length of the dataframe.
func (df *DataFrame) AddColumn(s *series.Series, inplace bool) *DataFrame {
	if s.Len() != df.Length() {
		panic(errors.CustomError("mismatched number of rows in the added series"))
	}

	modifiedDF := df
	if !inplace {
		modifiedDF = df.DeepCopy()
	}

	// append column
	modifiedDF.columns = append(modifiedDF.columns, s.GetColumn())

	// add the data
	modifiedDF.Data[s.GetColumn().Name] = s

	return modifiedDF
}
