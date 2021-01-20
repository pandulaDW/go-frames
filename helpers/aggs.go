package helpers

import (
	"github.com/pandulaDW/go-frames/base"
	"math"
)

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

// MaxSeries returns the maximum of a series
func MaxSeries(num []interface{}, dtype base.DType) interface{} {
	maxInt := math.MinInt64
	maxFloat := float32(math.MinInt64)

	switch dtype {
	case base.Int:
		for _, val := range num {
			if val.(int) > maxInt {
				maxInt = val.(int)
			}
		}
		return maxInt

	case base.Float:
		for _, val := range num {
			if val.(float32) > maxFloat {
				maxFloat = val.(float32)
			}
		}
		return maxFloat
	}
	return nil
}

// MinSeries returns the minimum of a series
func MinSeries(num []interface{}, dtype base.DType) interface{} {
	minInt := math.MaxInt64
	minFloat := float32(math.MaxInt64)

	switch dtype {
	case base.Int:
		for _, val := range num {
			if val.(int) < minInt {
				minInt = val.(int)
			}
		}
		return minInt

	case base.Float:
		for _, val := range num {
			if val.(float32) < minFloat {
				minFloat = val.(float32)
			}
		}
		return minFloat
	}
	return nil
}
