package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
)

// DataFrameData describes the shape of data stored in dataframe
type DataFrameData map[string]*series.Series

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	Data    DataFrameData
	Index   *series.Series
	length  int
	columns []*base.Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}
