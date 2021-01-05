package dataframes

// DType contains the supported data type definitions
type DType string

const (
	// Object is Text of mixed numeric values
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

// Column includes the Column name and type information
type Column struct {
	name  string
	dtype DType
}

// DataFrame includes the fields that describes a dataframe
type DataFrame struct {
	data    []interface{}
	length  int
	columns []Column
}

// Length returns the no of rows of the dataframe
func (df *DataFrame) Length() int {
	return df.length
}

// Columns returns the column names of the dataframe
func (df *DataFrame) Columns() []string {
	names := make([]string, len(df.columns))
	for i, val := range df.columns {
		names[i] = val.name
	}
	return names
}
