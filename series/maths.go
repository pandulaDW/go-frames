package series

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"strconv"
)

// Round will round the values of the series to the number of decimal places given.
// It will only work for a Float series, and will panic if a series with mismatched type is given
//
// Setting inplace to true will overwrite the current series and will return a new series otherwise
func (s *Series) Round(n int, inplace bool) *Series {
	if s.column.Dtype != base.Float {
		panic(errors.New("only a series with float type can be rounded"))
	}
	newData := make([]interface{}, 0, s.Len())
	format := "%." + fmt.Sprintf("%df", n)

	for i, val := range s.Data {
		floatVal, ok := val.(float64)
		if !ok {
			panic(errors.New(fmt.Sprintf("invalid value at row %d", i)))
		}
		roundedStr := fmt.Sprintf(format, floatVal)
		roundedVal, _ := strconv.ParseFloat(roundedStr, 64)

		if inplace {
			s.Data[i] = roundedVal
		} else {
			newData = append(newData, roundedVal)
		}
	}

	if len(newData) > 0 {
		return &Series{Data: newData, column: s.column}
	}

	return s
}
