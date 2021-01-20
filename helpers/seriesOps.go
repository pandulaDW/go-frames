package helpers

// ValueCounts will find the duplicate elements from a given array of
// empty interfaces and will return a map containing a value count for each unique elements
func ValueCounts(arr []interface{}) map[interface{}]int {
	valueCounts := map[interface{}]int{}

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if _, ok := valueCounts[val]; !ok {
			valueCounts[val] = 1
		} else {
			valueCounts[val]++
		}
	}

	return valueCounts
}

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
