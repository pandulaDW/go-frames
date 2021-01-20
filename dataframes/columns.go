package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/base"
)

// Column includes the Column name and type information
type Column struct {
	name     string
	dtype    base.DType
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
func (df *DataFrame) ColDType(colName string) (base.DType, error) {
	for _, val := range df.columns {
		if val.name == colName {
			return val.dtype, nil
		}
	}
	return "", errors.New("column not found")
}

// assertType would take a column of data as an argument and will infer the
// type of the column. If the type contains mix type data, it will default to Object type
func (df *DataFrame) assertType() {
	for i, col := range df.columns {
		for _, val := range df.Data[col.name] {
			switch val.(type) {
			case int64:
				df.columns[i].dtype = base.Int
			case float64:
				df.columns[i].dtype = base.Float
			case bool:
				df.columns[i].dtype = base.Bool
			default:
				df.columns[i].dtype = base.Object
			}
		}
	}
}
