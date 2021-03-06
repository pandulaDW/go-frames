package series

import "github.com/pandulaDW/go-frames/helpers"

// MemSize returns the memory size of the series
func (s *Series) MemSize() int {
	return helpers.GetRealSizeOf(s.Data)
}

// ValueCounts Return a map containing counts of unique values.
//
// The resulting map can then be converted to a dataframe and sorted
// to display the most frequently-occurring or least least-occurring element.
func (s *Series) ValueCounts() map[interface{}]int {
	countMap := make(map[interface{}]int)

	for _, val := range s.Data {
		_, ok := countMap[val]
		if !ok {
			countMap[val] = 1
		} else {
			countMap[val]++
		}
	}

	return countMap
}
