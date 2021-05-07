package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
)

// Apply will map each element of the series to the given function and
// will return a new series with a new column name "functionName(colName)"
func (s *Series) Apply(mapper base.ApplyFunc) (*Series, error) {
	seriesData := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		if val == nil {
			seriesData = append(seriesData, nil)
			continue
		}
		mappedVal, err := mapper(val)
		if err != nil {
			return nil, err
		}
		seriesData = append(seriesData, mappedVal)
	}

	colName := helpers.FunctionNameWrapper(helpers.GetFunctionName(mapper), s.column.Name)
	newSeries := NewSeries(colName, seriesData...)
	return newSeries, nil
}

func compose(operator, colName string, sBool ...*Series) *Series {
	initLen := sBool[0].Len()
	data := make([]interface{}, initLen)

	for _, series := range sBool {
		if series.Len() != initLen {
			err := fmt.Errorf("%s is invalid. %s", series.column.Name,
				errors.MismatchedNumOfRows(initLen, series.Len()).Error())
			panic(err)
		}
	}

	for i := 0; i < initLen; i++ {
		var val bool
		if operator == "AND" {
			val = true
		} else {
			val = false
		}

		for _, series := range sBool {
			sVal, ok := series.Data[i].(bool)
			if !ok {
				panic(errors.InvalidSeriesValError(series.Data[i], i, series.column.Name))
			}
			switch operator {
			case "AND":
				val = val && sVal
			case "OR":
				val = val || sVal
			}
		}
		data[i] = val
	}

	newS := &Series{Data: data, column: base.Column{Name: colName, Dtype: base.Bool}}
	return newS
}

// ComposeWithAnd takes in variadic number of base.Bool Series and return a single base.Bool Series
// by applying AND operation to all the provided Series.
//
// The function panics if Series with mismatched number of rows or wrongly typed Series are given.
func ComposeWithAnd(sBool ...*Series) *Series {
	return compose("AND", "compose-with-and", sBool...)
}

// ComposeWithOR takes in variadic number of base.Bool Series and return a single base.Bool Series
// by applying OR operation to all the provided Series.
//
// The function panics if Series with mismatched number of rows or wrongly typed Series are given.
func ComposeWithOR(sBool ...*Series) *Series {
	return compose("OR", "compose-with-or", sBool...)
}
