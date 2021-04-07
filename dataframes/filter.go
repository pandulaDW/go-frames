package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
)

// FilterBySeries takes in a Series of type base.Bool and will return a
// new DataFrame filtered by the true values.
//
// The function panics if the Series datatype is not base.Bool or if the Series
// length is not matching with the DataFrame length.
func (df *DataFrame) FilterBySeries(boolSeries *series.Series) *DataFrame {
	if boolSeries.GetColumn().Dtype != base.Bool {
		panic(errors.IncorrectDataType(base.Bool))
	}

	if df.length != boolSeries.Len() {
		panic(errors.MismatchedNumOfRows(df.length, boolSeries.Len()))
	}

	// append truth indices
	indices := make([]int, 0)
	for i, val := range boolSeries.Data {
		boolVal, ok := val.(bool)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, boolSeries.GetColumn().Name))
		}
		if boolVal {
			indices = append(indices, i)
		}
	}

	// create new series instances
	filteredSeriesArray := make([]*series.Series, 0, len(df.columns))

	for _, column := range df.columns {
		filteredData := make([]interface{}, 0, len(indices))
		for _, index := range indices {
			filteredData = append(filteredData, df.data[column.Name].Data[index])
		}
		filteredSeries := df.data[column.Name].ShallowCopy()
		filteredSeries.Data = filteredData
		filteredSeriesArray = append(filteredSeriesArray, filteredSeries)
	}

	return NewDataFrame(filteredSeriesArray...)
}

// Sample will return a new DataFrame with n number of rows randomly selected using the given seed value.
//
// If withReplacement is set to true, the sample can contain duplicate rows. And if withReplacement is set to
// false and if the length of the DataFrame is lower than n, then the returned sample will contain number of
// elements equal to the DataFrame length.
func (df *DataFrame) Sample(n uint64, seed int64, withReplacement bool) *DataFrame {
	seq := helpers.GenerateRandomSeries(n, uint64(df.length), seed, withReplacement)
	return df.ILoc(seq, df.Columns())
}
