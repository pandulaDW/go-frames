package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
)

// SetIndex sets the given column as the index. Panics if the column name is not found
func (df *DataFrame) SetIndex(colName string) *DataFrame {
	if helpers.LinearSearch(colName, helpers.ToInterfaceFromString(df.Columns())) == -1 {
		panic(errors.CustomError("column not found"))
	}

	df.Index = Index{Data: df.Data[colName], IsCustom: true}
	df.Drop(colName)

	return df
}
