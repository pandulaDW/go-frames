package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"time"
)

func helperCrud(s *Series, val interface{}, operation string) *Series {
	valS, isSeries := val.(*Series)
	if isSeries {
		if valS.Len() != s.Len() {
			panic(errors.MismatchedNumOfRows(s.Len(), valS.Len()))
		}
	}

	newS := s.ShallowCopy()
	data := make([]interface{}, s.Len())

	var curVal interface{}
	for i := range s.Data {
		if s.Data[i] == nil {
			data[i] = nil
		}
		if isSeries {
			curVal = valS.Data[i]
		} else {
			curVal = val
		}

		if s.column.Dtype == base.Int {
			intVal, ok := curVal.(int)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "int"))
			}
			sIntVal, ok := s.Data[i].(int)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = intVal + sIntVal
			case operation == "SUBTRACT":
				data[i] = sIntVal - intVal
			}
		}

		if s.column.Dtype == base.Float {
			floatVal, ok := curVal.(float64)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "float64"))
			}
			sFloatVal, ok := s.Data[i].(float64)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = floatVal + sFloatVal
			case operation == "SUBTRACT":
				data[i] = sFloatVal - floatVal
			}
		}

		if s.column.Dtype == base.Object {
			strVal, ok := curVal.(string)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "float64"))
			}
			sStringVal, ok := s.Data[i].(string)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = sStringVal + strVal
			}
		}

		if s.column.Dtype == base.DateTime {
			duration, ok := curVal.(time.Duration)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "time.Duration"))
			}
			sTVal, ok := s.Data[i].(time.Time)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = sTVal.Add(duration)
			}
		}
	}

	newS.Data = data
	return newS
}

func setOpFuncName(val interface{}, prefix string, s, newS *Series) {
	if valS, ok := val.(*Series); ok {
		newS.column.Name = fmt.Sprintf("%s(%s, %s)", prefix, s.column.Name, valS.column.Name)
	} else {
		newS.column.Name = fmt.Sprintf("%s(%s, %v)", prefix, s.column.Name, val)
	}
}

// Add will add the given value to each value in the calling Series and will return a new Series
// with added values.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be added with the corresponding value in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Add(val interface{}) *Series {
	newS := helperCrud(s, val, "ADD")
	setOpFuncName(val, "add", s, newS)
	return newS
}

// Subtract will subtract the given value from each value in the calling Series and will return a new Series
// with subtracted values.
//
// If another Series is passed as the val parameter using Col method, each value of the passed
// Series will be subtracted from the corresponding value in the calling Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Subtract(val interface{}) *Series {
	newS := helperCrud(s, val, "SUBTRACT")
	setOpFuncName(val, "subtract", s, newS)
	return newS
}