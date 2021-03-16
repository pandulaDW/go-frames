package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
)

// ApplyToRows applies a function along the rows of the DataFrame.
//
// The function returns a new DataFrame object and if an error encountered,
// it will return nil with the error.
func (df *DataFrame) ApplyToRows(fun base.ApplyFunc) (*DataFrame, error) {
	rowContent := make([][]string, 0)

	for i := 0; i < df.length; i++ {
		row := make([]string, 0)
		for _, column := range df.columns {
			currentVal := df.data[column.Name].Data[i]
			modifiedVal, err := fun(currentVal)
			if err != nil {
				return nil, err
			}
			strVal := fmt.Sprintf("%v", modifiedVal)
			row = append(row, strVal)
		}
		rowContent = append(rowContent, row)
	}

	modifiedDF := ConvertRowContentToDF(df.Columns(), rowContent)
	return modifiedDF, nil
}

// ApplyToColumns applies a function along the given set of columns of the DataFrame.
//
// The function returns a new DataFrame object and if an error encountered,
// it will return nil with the error.
func (df *DataFrame) ApplyToColumns(cols []string, fun base.ApplyFunc) (*DataFrame, error) {
	copiedDF := df.ShallowCopy()

	for _, col := range cols {
		s, ok := copiedDF.data[col]
		if !ok {
			return nil, errors.ColumnNotFound(col)
		}
		result, err := s.Apply(fun)
		if err != nil {
			return nil, err
		}
		result.SetColName(col)
		copiedDF.data[col] = result
	}

	return copiedDF, nil
}
