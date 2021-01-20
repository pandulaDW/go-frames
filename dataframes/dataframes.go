package dataframes

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
