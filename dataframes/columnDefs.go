package dataframes

import "errors"

// Column includes the Column name and type information
type Column struct {
	name     string
	dtype    DType
	colIndex int
}

// Columns returns the column names of the dataframe
func (df *DataFrame) Columns() []string {
	names := make([]string, len(df.columns))
	for i, val := range df.columns {
		names[i] = val.name
	}
	return names
}

// ColDType returns the datatype of the column. If the column is not
// found, it will return an error with an empty string
func (df *DataFrame) ColDType(colName string) (DType, error) {
	for _, val := range df.columns {
		if val.name == colName {
			return val.dtype, nil
		}
	}
	return "", errors.New("Column not found")
}
