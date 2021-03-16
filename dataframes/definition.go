package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
)

// DataFrameData describes the shape of data stored in dataframe
type DataFrameData map[string]*series.Series

// Index describes the index of a dataframe
type Index struct {
	Data     *series.Series
	IsCustom bool
}

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	data    DataFrameData
	Index   Index
	length  int
	columns []*base.Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}
