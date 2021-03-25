package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"strconv"
)

// Round will round the values of the series to the number of decimal places given.
// It will only work for a Float series, and will panic if a series with mismatched type is given
//
// The function will return a new series if rounding is successful.
func (s *Series) Round(n int) *Series {
	if s.column.Dtype != base.Float {
		panic(errors.IncorrectDataType(base.Float))
	}
	newData := make([]interface{}, 0, s.Len())
	format := "%." + fmt.Sprintf("%df", n)

	for i, val := range s.Data {
		if val == nil {
			newData = append(newData, nil)
			continue
		}
		floatVal, ok := val.(float64)
		if !ok {
			panic(errors.InvalidRowValue(i))
		}
		roundedStr := fmt.Sprintf(format, floatVal)
		roundedVal, _ := strconv.ParseFloat(roundedStr, 64)
		newData = append(newData, roundedVal)
	}

	rounded := &Series{Data: newData, column: s.column}
	rounded.column.Name = helpers.FunctionNameWrapper("round", s.column.Name)
	return rounded
}
