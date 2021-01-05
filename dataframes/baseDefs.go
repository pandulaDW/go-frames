package dataframes

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
	// NA represents empty cells
	NA DType = "NA"
)

// DataFrameData describes the shape of data stored in dataframe
type DataFrameData map[string][]interface{}

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	data    DataFrameData
	length  int
	columns []Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}

// CreateDataFrame creates a dataframe using given data
func CreateDataFrame(data [][]interface{}, columns []string) *DataFrame {
	df := new(DataFrame)
	df.data = make(DataFrameData)
	df.columns = make([]Column, 0)

	for i, colData := range data {
		df.data[columns[i]] = colData
		column := Column{}
		column.name = columns[i]
		column.dtype = Object // temporary
		column.colIndex = i
		df.columns = append(df.columns, column)
	}

	return df
}
