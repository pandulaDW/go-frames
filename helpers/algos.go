package helpers

// LinearSearch would search for the element in the given array and if the element is found
// it will return the index of the first occurrence in which it is found.
// If the element is not present it will return -1.
//
// Use ToInterfaceFrom functions to quickly convert slices into interface slices to be used by the function.
//
// Panics, if non-comparable elements are found
func LinearSearch(element interface{}, arr []interface{}) int {
	for i, val := range arr {
		if element == val {
			return i
		}
	}
	return -1
}
