package series

type funcMapType func(val interface{}) (interface{}, error)

// Map will map each element of the series to the given function and
// will return a new series
func (s *Series) Map(mapper funcMapType) ([]interface{}, error) {
	newSeries := make([]interface{}, 0, s.Len())
	for _, val := range s.Data {
		mappedVal, err := mapper(val)
		if err != nil {
			return nil, err
		}
		newSeries = append(newSeries, mappedVal)
	}
	return newSeries, nil
}
