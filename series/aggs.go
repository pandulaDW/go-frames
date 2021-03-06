package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"math"
	"time"
)

// Max returns the maximum value of the series based on it's data type.
// Returns nil, when the dtype is not applicable
func (s *Series) Max() interface{} {
	maxInt := math.MinInt64
	maxFloat := float64(math.MinInt64)

	switch s.column.Dtype {
	case base.Int:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			intVal, ok := val.(int)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			if intVal > maxInt {
				maxInt = intVal
			}
		}
		return maxInt
	case base.Float:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			floatVal, ok := val.(float64)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			if floatVal > maxFloat {
				maxFloat = floatVal
			}
		}
		return maxFloat
	case base.DateTime:
		maxDateTime, ok := s.Data[0].(time.Time)
		if !ok {
			panic(errors.InvalidSeriesValError(s.Data[0], 0, s.column.Name))
		}
		for i, val := range s.Data[1:] {
			if val == nil {
				continue
			}
			parsedVal, ok := val.(time.Time)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i+1, s.column.Name))
			}
			if maxDateTime.Before(parsedVal) {
				maxDateTime = parsedVal
			}
		}
		return maxDateTime
	default:
		return nil
	}
}

// Min returns the minimum value of the series based on it's data type.
// Returns nil, when the dtype is not applicable
func (s *Series) Min() interface{} {
	minInt := math.MaxInt64
	minFloat := float64(math.MaxInt64)

	switch s.column.Dtype {
	case base.Int:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			intVal, ok := val.(int)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			if intVal < minInt {
				minInt = intVal
			}
		}
		return minInt
	case base.Float:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			floatVal, ok := val.(float64)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			if floatVal < minFloat {
				minFloat = floatVal
			}
		}
		return minFloat
	case base.DateTime:
		minDataTime, ok := s.Data[0].(time.Time)
		if !ok {
			panic(errors.InvalidSeriesValError(s.Data[0], 0, s.column.Name))
		}
		for i, val := range s.Data[1:] {
			if val == nil {
				continue
			}
			parsedTime, ok := val.(time.Time)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i+1, s.column.Name))
			}
			if minDataTime.After(parsedTime) {
				minDataTime = parsedTime
			}
		}
		return minDataTime
	default:
		return nil
	}
}

// Sum returns the total value of the series for integer and floating type
// as a float64. Panic when Dtype is not applicable
func (s *Series) Sum() float64 {
	var sumInt int
	var sumFloat float64

	switch s.column.Dtype {
	case base.Int:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			intVal, ok := val.(int)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			sumInt += intVal
		}
		return float64(sumInt)
	case base.Float:
		for i, val := range s.Data {
			if val == nil {
				continue
			}
			floatVal, ok := val.(float64)
			if !ok {
				panic(errors.InvalidSeriesValError(val, i, s.column.Name))
			}
			sumFloat += floatVal
		}
		return sumFloat
	default:
		panic(errors.CustomError("sum can only be applied for a numerical series"))
	}
}

// Avg returns the average value of the series for integer and floating type.
// Returns nil, when the dtype is not applicable
func (s *Series) Avg() float64 {
	countWithoutNA := s.Len() - s.CountOfNA()
	avgVal := s.Sum() / float64(countWithoutNA)
	return avgVal
}
