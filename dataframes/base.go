package dataframes

// DType contains the supported data type definitions
type DType string

const (
	// Object is Text or mixed numeric values
	Object DType = "Object"
	// Int is int64 typed numeric values
	Int DType = "Int"
	// Float is Float64 typed floating point values
	Float DType = "Float"
	// Bool is True/False values
	Bool DType = "Bool"
	// DateTime is Date and Time values
	DateTime DType = "DateTime"
	// NA represents empty cells
	NA DType = "NA"
)

// DataFrameData describes the shape of data stored in dataframe
type DataFrameData map[string][]interface{}

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	Data    DataFrameData
	length  int
	columns []Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}
