package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
)

// CreateDataFrame creates a dataframe using given parameters of series.
// if the length of the series are mismatching, it will panic with an error
func CreateDataFrame(data ...*series.Series) *DataFrame {
	df := new(DataFrame)

	if len(data) == 0 {
		return df
	}

	df.Data = make(DataFrameData)
	df.columns = make([]*base.Column, 0)
	df.length = data[0].Len()

	// Populate the dataframe and the columns
	for i, s := range data {
		if s.Len() != df.length {
			panic("Mismatched row lengths found. Dataframe can only contain equal number of rows")
		}
		df.Data[s.GetColumn().Name] = s
		s.SetColIndex(i)
		df.columns = append(df.columns, s.GetColumn())
	}

	// infer the types
	df.assertType()

	return df
}
