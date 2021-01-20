package helpers

import "math"

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

// MaxInt returns the maximum of two integers
func MaxInt(num1, num2 int64) int64 {
	if num1 > num2 {
		return num1
	}
	return num2
}

// MaxFloat returns the maximum of two floats
func MaxFloat(num1, num2 float64) float64 {
	if num1 > num2 {
		return num1
	}
	return num2
}
