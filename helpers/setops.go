package helpers

// Difference returns the set difference of two slices with any given values.
//
// Only use acceptable key types for a Go map
func Difference(a, b []interface{}) []interface{} {
	m := make(map[interface{}]int)
	diff := make([]interface{}, 0)

	populateMap := func(arr []interface{}) {
		for _, val := range arr {
			_, ok := m[val]
			if !ok {
				m[val] = 0
			}
		}
	}

	populateMap(a)
	populateMap(b)

	for key := range m {
		diff = append(diff, key)
	}

	return diff
}
