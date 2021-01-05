package dataframes

import (
	orderedmap "github.com/wk8/go-ordered-map"
)

// DType contains the supported data type definitions
type DType string

const (
	// Object is Text or mixed numeric values
	Object DType = "Object"
	// Int64 is int64 typed numeric values
	Int64 DType = "Int64"
	// Float64 is Float64 typed floating point values
	Float64 DType = "Float64"
	// Bool is True/False values
	Bool DType = "Bool"
	// DateTime is Date and Time values
	DateTime DType = "DateTime"
	// Category is finite list of text values
	Category DType = "Category"
)

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	data    *orderedmap.OrderedMap
	length  int
	columns []Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}

// CreateDataFrame creates a dataframe using given data
func CreateDataFrame(data *orderedmap.OrderedMap) *DataFrame {
	df := new(DataFrame)
	df.data = data

	index := 0
	for pair := data.Oldest(); pair != nil; pair = pair.Next() {
		column := Column{}
		column.name = pair.Key.(string)
		column.dtype = Object // temporary
		column.colIndex = index
		index++
	}

	return df
}
