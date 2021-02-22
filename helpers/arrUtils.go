package helpers

import "math"

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

// MaxIntSlice returns the maximum element of a given int array
func MaxIntSlice(nums []int) int {
	max := math.MinInt32
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	return max
}

// ReverseArray reverses the given slice
func ReverseArray(arr []interface{}) []interface{} {
	length := len(arr)
	reversedArr := make([]interface{}, length)

	for i := length; i > 0; i-- {
		reversedArr[length-i] = arr[i-1]
	}

	return reversedArr
}
