package dataframes

import (
	"fmt"
)

// ApplyFunc takes in any value and return another value alongside an error if encountered
type ApplyFunc func(val interface{}) (interface{}, error)

// ApplyToRows applies a function along the rows of the DataFrame.
//
// The function returns the current DataFrame object and if an error encountered,
// it will return nil with the error.
func (df *DataFrame) ApplyToRows(fun ApplyFunc) (*DataFrame, error) {
	rowContent := make([][]string, 0)

	for i := 0; i < df.length; i++ {
		row := make([]string, 0)
		for _, column := range df.columns {
			currentVal := df.Data[column.Name].Data[i]
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
