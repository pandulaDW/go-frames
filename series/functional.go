package series

type funcMapType func(val interface{}) (interface{}, error)

// Map will map each element of the series to the given function and
// will return a new series
func Map(series []interface{}, mapper funcMapType) ([]interface{}, error) {
	newSeries := make([]interface{}, 0, len(series))
	for _, val := range series {
		mappedVal, err := mapper(val)
		if err != nil {
			return nil, err
		}
		newSeries = append(newSeries, mappedVal)
	}
	return newSeries, nil
}
