package dataframes

import "errors"

// Max returns the maximum of the column specified. Panics if
// the column is not found
func (df *DataFrame) Max(col string) interface{} {
	var series []interface{}
	var dtype *DType

	for _, val := range df.columns {
		if val.name == col {
			series = df.Data[col]
			dtype = &val.dtype
		}
	}

	if dtype == nil {
		panic(errors.New("column not found"))
	}

	var max interface{}

	for _, val := range series {
		switch *dtype {
		case Int:
			if val == 5 {
			}
		}
	}

	return max
}
