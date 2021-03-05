package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/helpers"
	"time"
)

// GetMaxLength returns the column length of the series. This can be useful for
// displaying dataframes and series.
func (s *Series) GetMaxLength() int {
	strLengths := make([]int, 0, s.Len())

	for _, val := range s.Data {
		if s.column.Dtype == base.DateTime && helpers.IsTimeSet(val.(time.Time)) {
			strLengths = append(strLengths, 11)
			continue
		}
		strRepr := fmt.Sprintf("%v", val)
		strLengths = append(strLengths, len(strRepr))
	}
	strLengths = append(strLengths, len(s.column.Name))

	return helpers.MaxIntSlice(strLengths)
}
