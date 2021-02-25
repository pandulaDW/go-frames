package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
)

// NewDataFrame creates a dataframe using given parameters of series.
//
// The function panics if the length of the series are mismatching or a duplicated column name
// is provided.
func NewDataFrame(data ...*series.Series) *DataFrame {
	df := new(DataFrame)

	if len(data) == 0 {
		return df
	}

	// set the variables
	df.Data = make(DataFrameData)
	df.columns = make([]*base.Column, 0)
	df.length = data[0].Len()

	// set the index
	indices := helpers.ToInterfaceFromInt(helpers.Range(0, df.length, 1))
	df.Index = Index{
		Data:     series.NewSeries("#", indices...),
		IsCustom: false,
	}

	// Populate the dataframe and the columns
	for i, s := range data {
		colName := s.GetColumn().Name
		if s.Len() != df.length {
			panic(errors.MismatchedNumOfRows(df.length, s.Len()))
		}
		if _, ok := df.Data[colName]; ok {
			panic(errors.DuplicatedColumn(colName))
		}
		df.Data[colName] = s
		s.SetColIndex(i)
		df.columns = append(df.columns, s.GetColumn())
	}

	return df
}

// DeepCopy will create a new dataframe and will return a pointer to the new dataframe.
//
// As the underlying data is also copied, can cause memory leak in large dataframes.
func (df *DataFrame) DeepCopy() *DataFrame {
	copiedSeriesArr := make([]*series.Series, 0, len(df.columns))
	for _, col := range df.columns {
		copiedSeriesArr = append(copiedSeriesArr, df.Data[col.Name].DeepCopy())
	}
	return NewDataFrame(copiedSeriesArr...)
}

// ShallowCopy will create a new dataframe and will return a pointer to the new dataframe.
//
// Underlying series objects will be kept the same.
func (df *DataFrame) ShallowCopy() *DataFrame {
	seriesArr := make([]*series.Series, 0, len(df.columns))
	for _, col := range df.columns {
		seriesArr = append(seriesArr, df.Data[col.Name])
	}
	return NewDataFrame(seriesArr...)
}
