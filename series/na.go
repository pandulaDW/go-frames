package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/helpers"
)

// Return a boolean same-sized Series indicating if the values are NA.
//
// NA values such as N/A, None, empty strings, gets mapped to True values.
// Everything else gets mapped to False values.
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
