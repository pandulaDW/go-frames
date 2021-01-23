package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
)

// NewDataFrame creates a dataframe using given parameters of series.
// if the length of the series are mismatching, it will panic with an error
func NewDataFrame(data ...*series.Series) *DataFrame {
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
			panic(errors.New("mismatched row lengths found. " +
				"Dataframe can only contain equal number of rows"))
		}
		df.Data[s.GetColumn().Name] = s
		s.SetColIndex(i)
		df.columns = append(df.columns, s.GetColumn())
	}

	return df
}
