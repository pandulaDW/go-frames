package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"strings"
)

type stringMethod func(val string) string

// helperStringMethods applies the function provided and returns a new Series
func helperStringMethods(s *Series, fun stringMethod) *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	data := make([]interface{}, s.Len())

	for i, val := range s.Data {
		strVal, ok := val.(string)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}
		data[i] = fun(strVal)
	}

	return NewSeries("test", data...)
}

// Lower will return a new series with the values lowercased.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Lower() *Series {
	return helperStringMethods(s, strings.ToLower)
}

// Upper will return a new series with the values uppercased.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Upper() *Series {
	return helperStringMethods(s, strings.ToUpper)
}

// Capitalized will return a new series Title returns a copy of the strings
// with all Unicode letters that begin words mapped to their Unicode title case.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Capitalized() *Series {
	return helperStringMethods(s, strings.Title)
}
