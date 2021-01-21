package series

import (
	"github.com/pandulaDW/go-frames/base"
	"math"
)

// Max returns the maximum value of the series based on it's data type
func (s *Series) Max() interface{} {
	maxInt := math.MinInt64
	maxFloat := float64(math.MinInt64)

	switch s.column.Dtype {
	case base.Int:
		for _, val := range s.Data {
			if val.(int) > maxInt {
				maxInt = val.(int)
			}
		}
		return maxInt
	case base.Float:
		for _, val := range s.Data {
			if val.(float64) > maxFloat {
				maxFloat = val.(float64)
			}
		}
		return maxFloat
	}
	return nil
}

// Min returns the minimum value of the series based on it's data type
func (s *Series) Min() interface{} {
	minInt := math.MaxInt64
	minFloat := float64(math.MaxInt64)

	switch s.column.Dtype {
	case base.Int:
		for _, val := range s.Data {
			if val.(int) < minInt {
				minInt = val.(int)
			}
		}
		return minInt
	case base.Float:
		for _, val := range s.Data {
			if val.(float64) < minFloat {
				minFloat = val.(float64)
			}
		}
		return minFloat
	}
	return nil
}

// Sum returns the total value of the series for integer and floating type
// as a float64
func (s *Series) Sum() float64 {
	var sumInt int
	var sumFloat float64

	switch s.column.Dtype {
	case base.Int:
		for _, val := range s.Data {
			sumInt += val.(int)
		}
		return float64(sumInt)
	case base.Float:
		for _, val := range s.Data {
			sumFloat += val.(float64)
		}
		return sumFloat
	default:
		panic("Sum can only be applied for numerical series")
	}
}

// Avg returns the average value of the series for integer and floating type
func (s *Series) Avg() float64 {
	return s.Sum() / float64(s.Len())
}
