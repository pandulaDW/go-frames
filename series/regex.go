package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"regexp"
)

// RegexExtract returns a string holding the text of the leftmost
// match in s of the given regular expression for each value of the Series.
//
// If there is no match or if the value is nil or an empty string, the Series val
// will be an empty string
//
// The function panics if the Series datatype is not base.Object.
func (s *Series) RegexExtract(re *regexp.Regexp) *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	extractedData := make([]interface{}, 0, s.Len())

	for i, val := range s.Data {
		if val == nil {
			extractedData = append(extractedData, "")
			continue
		}
		strVal, ok := val.(string)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}

		findString := re.FindString(strVal)
		extractedData = append(extractedData, findString)
	}

	newS := s.ShallowCopy()
	newS.Data = extractedData
	newS.column.Name = helpers.FunctionNameWrapper("regex_extract", s.column.Name)

	return newS
}
