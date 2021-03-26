package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"regexp"
)

// RegexContains returns a boolean Series indicating whether the string value
// contains any match of the regular expression re for each value in the Series.
//
// If there is no match or if the value is nil or an empty string, the Series val
// will return false.
//
// The function panics if the Series datatype is not base.Object.
func (s *Series) RegexContains(re *regexp.Regexp) *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	extractedData := make([]interface{}, 0, s.Len())

	for i, val := range s.Data {
		if val == nil {
			extractedData = append(extractedData, false)
			continue
		}
		strVal, ok := val.(string)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}

		match := re.MatchString(strVal)
		extractedData = append(extractedData, match)
	}

	newS := s.ShallowCopy()
	newS.Data = extractedData
	newS.column.Name = helpers.FunctionNameWrapper("regex_extract", s.column.Name)
	newS.column.Dtype = base.Bool

	return newS
}
