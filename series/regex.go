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
	newS.column.Name = helpers.FunctionNameWrapper("regex_contains", s.column.Name)
	newS.column.Dtype = base.Bool

	return newS
}

// RegexExtract returns a Series where each returned value of the Series is the submatch corresponding
// to the index value given. If a submatch is not present or value is nil/empty, then the returned value will be
// an empty string.
//
// Submatches are matches of parenthesized subexpressions (also known as capturing groups) within the
// regular expression, numbered from left to right in order of opening parenthesis. Submatch 0 is the match
// of the entire expression, submatch 1 the match of the first parenthesized subexpression, and so on.
//
// The function panics if the Series datatype is not base.Object or if the index is out of range.
func (s *Series) RegexExtract(re *regexp.Regexp, index int) *Series {
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

		matches := re.FindStringSubmatch(strVal)
		if matches == nil {
			extractedData = append(extractedData, "")
			continue
		}

		if index > len(matches) {
			panic(errors.CustomError("index is out of range"))
		}

		extractedData = append(extractedData, matches[index])
	}

	newS := s.ShallowCopy()
	newS.Data = extractedData
	newS.column.Name = helpers.FunctionNameWrapper("regex_extract", s.column.Name)

	return newS
}

// RegexReplace returns a Series where each returned value of the Series is a copy of it's src value,
// replacing matches of the Regexp with the replacement string replaceStr.
//
// Inside replaceStr, $ signs are interpreted as in regexp.Expand, so for instance $1 represents
// the text of the first submatch.
//
// The function panics if the Series datatype is not base.Object.
func (s *Series) RegexReplace(re *regexp.Regexp, replaceStr string) *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	replacedData := make([]interface{}, 0, s.Len())

	for i, val := range s.Data {
		if val == nil {
			replacedData = append(replacedData, "")
			continue
		}
		strVal, ok := val.(string)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}

		replacedVal := re.ReplaceAllString(strVal, replaceStr)
		replacedData = append(replacedData, replacedVal)
	}

	newS := s.ShallowCopy()
	newS.Data = replacedData
	newS.column.Name = helpers.FunctionNameWrapper("regex_replace", s.column.Name)

	return newS
}
