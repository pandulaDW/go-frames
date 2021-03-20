package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"strings"
)

type stringMethod func(val string) string
type stringBoolMethod func(val1, val2 string) bool

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

// helperStringBooleanMethods applies the function provided and returns a new Series
func helperStringBooleanMethods(s *Series, str string, fun stringBoolMethod) *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	data := make([]interface{}, s.Len())

	for i, val := range s.Data {
		strVal, ok := val.(string)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}
		data[i] = fun(strVal, str)
	}

	return NewSeries("test", data...)
}

// Lower will return a new Series with the values lowercased.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Lower() *Series {
	return helperStringMethods(s, strings.ToLower)
}

// Upper will return a new Series with the values uppercased.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Upper() *Series {
	return helperStringMethods(s, strings.ToUpper)
}

// Capitalized will return a new Series where all Unicode letters that
// begin words mapped to their Unicode title case.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Capitalized() *Series {
	return helperStringMethods(s, strings.Title)
}

// Trim returns a new Series where for each value, all leading and trailing white
// space removed, as defined by Unicode.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Trim() *Series {
	return helperStringMethods(s, strings.TrimSpace)
}

// Contains returns a new Series of base.Bool type which reports whether substr is within each
// value of the Series.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) Contains(substr string) *Series {
	return helperStringBooleanMethods(s, substr, strings.Contains)
}

// StartsWith returns a new Series of base.Bool type which reports whether the value begins with
// the given prefix for each value of the Series.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) StartsWith(prefix string) *Series {
	return helperStringBooleanMethods(s, prefix, strings.HasPrefix)
}

// EndsWith returns a new Series of base.Bool type which reports whether the value ends with
// the given suffix for each value of the Series.
//
// The function panics if the series data type of the series is not base.Object.
func (s *Series) EndsWith(prefix string) *Series {
	return helperStringBooleanMethods(s, prefix, strings.HasSuffix)
}
