package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/helpers"
)

// Apply will map each element of the series to the given function and
// will return a new series with a new column name "functionName(colName)"
func (s *Series) Apply(mapper base.ApplyFunc) (*Series, error) {
	seriesData := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		mappedVal, err := mapper(val)
		if err != nil {
			return nil, err
		}
		seriesData = append(seriesData, mappedVal)
	}

	colName := helpers.FunctionNameWrapper(helpers.GetFunctionName(mapper), s.column.Name)
	newSeries := NewSeries(colName, seriesData...)
	return newSeries, nil
}
