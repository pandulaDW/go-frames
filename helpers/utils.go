package helpers

import (
	"math"
)

// Max returns the maximum element of a given int array
func Max(nums []int) int {
	max := math.MinInt32
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	return max
}

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
