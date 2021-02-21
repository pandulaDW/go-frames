package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/errors"
)

// Loc will return a slice of data indexed by an integer array.
//
// Panics if out of range indices are found.
//
// helpers.Range() function can be used to return an indices range.
func (s *Series) Loc(indices []int) []interface{} {
	dataSlice := make([]interface{}, 0, len(indices))

	for _, index := range indices {
		if index < 0 || index > s.Len() {
			panic(errors.CustomError(fmt.Sprintf("index %d is out of range", index)))
		}
		dataSlice = append(dataSlice, s.Data[index])
	}

	return dataSlice
}
