package series

import "github.com/pandulaDW/go-frames/helpers"

// MemSize returns the memory size of the series
func (s *Series) MemSize() int {
	return helpers.GetRealSizeOf(s.Data)
}

// ValueCounts Returns a map containing counts of unique values.
//
// The resulting map can then be converted to a dataframe and sorted
// to display the most frequently-occurring or least least-occurring element.
func (s *Series) ValueCounts() map[interface{}]interface{} {
	countMap := make(map[interface{}]interface{})

	for _, val := range s.Data {
		_, ok := countMap[val]
		if !ok {
			countMap[val] = 1
		} else {
			current := countMap[val]
			countMap[val] = current.(int) + 1
		}
	}

	return countMap
}
