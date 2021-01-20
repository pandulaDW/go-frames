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
	maxInt := int64(math.MinInt64)
	maxFloat := float64(math.MinInt64)

	switch dtype {
	case base.Int:
		for _, val := range num {
			if val.(int64) > maxInt {
				maxInt = val.(int64)
			}
		}
		return maxInt

	case base.Float:
		for _, val := range num {
			if val.(float64) > maxFloat {
				maxFloat = val.(float64)
			}
		}
		return maxFloat
	}
	return nil
}

// MinSeries returns the minimum of a series
func MinSeries(num []interface{}, dtype base.DType) interface{} {
	minInt := int64(math.MaxInt64)
	minFloat := float64(math.MaxInt64)

	switch dtype {
	case base.Int:
		for _, val := range num {
			if val.(int64) < minInt {
				minInt = val.(int64)
			}
		}
		return minInt

	case base.Float:
		for _, val := range num {
			if val.(float64) < minFloat {
				minFloat = val.(float64)
			}
		}
		return minFloat
	}
	return nil
}
