package series

import (
	"github.com/pandulaDW/go-frames/base"
	"math"
)

// Max returns the maximum value of the series based on it's data type
func (s *Series) Max() interface{} {
	maxInt := math.MinInt64
	maxFloat := float32(math.MinInt64)

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
			if val.(float32) > maxFloat {
				maxFloat = val.(float32)
			}
		}
		return maxFloat
	}
	return nil
}

// Min returns the minimum value of the series based on it's data type
func (s *Series) Min() interface{} {
	minInt := math.MaxInt64
	minFloat := float32(math.MaxInt64)

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
			if val.(float32) < minFloat {
				minFloat = val.(float32)
			}
		}
		return minFloat
	}
	return nil
}
