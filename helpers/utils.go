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
