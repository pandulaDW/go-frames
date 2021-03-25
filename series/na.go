package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
)

// IsNa returns a boolean same-sized Series indicating if the values are NA.
//
// NA values such as N/A, None, empty strings, gets mapped to true values.
// Everything else gets mapped to false values.
//
// This is the Boolean inverse of NotNa.
func (s *Series) IsNa() *Series {
	data := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		if val == nil {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}

	boolS := s.ShallowCopy()
	boolS.Data = data
	boolS.column.Dtype = base.Bool
	boolS.column.ColIndex = 0
	boolS.column.Name = helpers.FunctionNameWrapper("isna", s.column.Name)

	return boolS
}

// NotNa returns a boolean same-sized Series indicating if the values are not NA.
//
// NA values such as N/A, None, empty strings, gets mapped to false values.
// Everything else gets mapped to true values.
func (s *Series) NotNa() *Series {
	data := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		if val != nil {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}

	boolS := s.ShallowCopy()
	boolS.Data = data
	boolS.column.Dtype = base.Bool
	boolS.column.ColIndex = 0
	boolS.column.Name = helpers.FunctionNameWrapper("notna", s.column.Name)

	return boolS
}

// CountOfNA returns the count of NA values in a given Series.
func (s *Series) CountOfNA() int {
	count := 0
	for _, val := range s.Data {
		if val == nil {
			count++
		}
	}
	return count
}

// IsBlank returns a boolean same-sized Series indicating if the values are blank.
//
// The function panics if the Series datatype is not base.Object
func (s *Series) IsBlank() *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	data := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		if val == "" {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}

	boolS := s.ShallowCopy()
	boolS.Data = data
	boolS.column.Dtype = base.Bool
	boolS.column.ColIndex = 0
	boolS.column.Name = helpers.FunctionNameWrapper("isBlank", s.column.Name)

	return boolS
}

// NotBlank returns a boolean same-sized Series indicating if the values are not blank.
//
// The function panics if the Series datatype is not base.Object
func (s *Series) NotBlank() *Series {
	if s.column.Dtype != base.Object {
		panic(errors.IncorrectDataType(base.Object))
	}

	data := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		if val != "" {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}

	boolS := s.ShallowCopy()
	boolS.Data = data
	boolS.column.Dtype = base.Bool
	boolS.column.ColIndex = 0
	boolS.column.Name = helpers.FunctionNameWrapper("notBlank", s.column.Name)

	return boolS
}
