package helpers

// ToIntArray converts an array of empty interfaces to an int array
func ToIntArray(arr []interface{}) []int {
	intArray := make([]int, 0, len(arr))

	for _, val := range arr {
		intArray = append(intArray, val.(int))
	}

	return intArray
}

// ToInterfaceFromInt converts an array of int array to an empty interface array
func ToInterfaceFromInt(arr []int) []interface{} {
	interfaceArray := make([]interface{}, 0, len(arr))

	for _, val := range arr {
		interfaceArray = append(interfaceArray, val)
	}

	return interfaceArray
}

// ToInterfaceFromString converts an array of string array to an empty interface array
func ToInterfaceFromString(arr []string) []interface{} {
	interfaceArray := make([]interface{}, 0, len(arr))

	for _, val := range arr {
		interfaceArray = append(interfaceArray, val)
	}

	return interfaceArray
}
