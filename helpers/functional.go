package helpers

// Predicate is the definition of a function which takes in any element and returns
// a boolean
type Predicate func(element interface{}) (bool, error)

// Filter method creates an returns a new slice with all elements that pass the test implemented
// by the provided function.
//
// If an error is encountered, a nil value will be returned as a slice along with the error
func Filter(arr []interface{}, test Predicate) ([]interface{}, error) {
	filteredArr := make([]interface{}, 0)

	for _, val := range arr {
		ok, err := test(val)
		if err != nil {
			return nil, err
		}
		if ok {
			filteredArr = append(filteredArr, val)
		}
	}

	return filteredArr, nil
}
