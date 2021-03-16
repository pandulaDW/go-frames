package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
)

// SetIndex sets the given column as the index. Panics if the column name is not found
func (df *DataFrame) SetIndex(colName string) *DataFrame {
	if helpers.LinearSearch(colName, helpers.ToInterfaceFromString(df.Columns())) == -1 {
		panic(errors.ColumnNotFound(colName))
	}

	df.Index = Index{Data: df.data[colName], IsCustom: true}
	df.Drop(colName)

	return df
}

// SetIndexBySeries sets the index column as the given series. The function panics if the
// length of the series is different from dataframe length
func (df *DataFrame) SetIndexBySeries(s *series.Series) *DataFrame {
	if s.Len() != df.length {
		panic(errors.MismatchedNumOfRows(df.length, s.Len()))
	}

	df.Index = Index{Data: s, IsCustom: true}
	return df
}

// ResetIndex will Reset the index of the DataFrame, and use the default one instead.
// If the DataFrame still uses the default index, no change will be made.
//
// If drop is set to true, it will drop the current index column and if false, current index column column
// will be part of the dataframe
func (df *DataFrame) ResetIndex(drop bool) *DataFrame {
	if !df.Index.IsCustom {
		return df
	}

	if !drop {
		currentIndex := df.Index.Data
		df.WithColumn(currentIndex)
	}

	indices := helpers.ToInterfaceFromInt(helpers.Range(0, df.length, 1))
	df.Index = Index{
		Data:     series.NewSeries("#", indices...),
		IsCustom: false,
	}

	return df
}
