package series

type funcMapType func(val interface{}) (interface{}, error)

// Map will map each element of the series to the given function and
// will return a new series
func (s *Series) Map(mapper funcMapType) (*Series, error) {
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
