package series

import "github.com/pandulaDW/go-frames/base"

// Apply will map each element of the series to the given function and
// will return a new series
func (s *Series) Apply(mapper base.ApplyFunc) (*Series, error) {
	seriesData := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		mappedVal, err := mapper(val)
		if err != nil {
			return nil, err
		}
		seriesData = append(seriesData, mappedVal)
	}

	newSeries := Series{column: s.column, Data: seriesData}
	return &newSeries, nil
}
