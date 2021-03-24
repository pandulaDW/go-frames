package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
)

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
