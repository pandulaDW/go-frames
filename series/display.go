package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/helpers"
)

// GetMaxLength returns the column length of the series. This can be useful for
// displaying dataframes and series.
func (s *Series) GetMaxLength() int {
	strLengths := make([]int, 0, s.Len())

	for _, val := range s.Data {
		strRepr := fmt.Sprintf("%v", val)
		strLengths = append(strLengths, len(strRepr))
	}
	strLengths = append(strLengths, len(s.column.Name))

	return helpers.MaxIntSlice(strLengths)
}
