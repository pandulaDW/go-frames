package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"time"
)

func helperCrud(s *Series, val interface{}, operation string, conditional bool) *Series {
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
		if s.Data[i] == nil && !conditional {
			data[i] = nil
			continue
		}
		if s.Data[i] == nil && conditional {
			data[i] = false
			continue
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
			case operation == "GT":
				data[i] = sIntVal > intVal
			case operation == "GTE":
				data[i] = sIntVal >= intVal
			case operation == "LT":
				data[i] = sIntVal < intVal
			case operation == "LTE":
				data[i] = sIntVal <= intVal
			case operation == "EQ":
				data[i] = sIntVal == intVal
			case operation == "NEQ":
				data[i] = sIntVal != intVal
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
			case operation == "GT":
				data[i] = sFloatVal > floatVal
			case operation == "GTE":
				data[i] = sFloatVal >= floatVal
			case operation == "LT":
				data[i] = sFloatVal < floatVal
			case operation == "LTE":
				data[i] = sFloatVal <= floatVal
			case operation == "EQ":
				data[i] = sFloatVal == floatVal
			case operation == "NEQ":
				data[i] = sFloatVal != floatVal
			}
		}

		if s.column.Dtype == base.Object {
			strVal, ok := curVal.(string)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "string"))
			}
			sStringVal, ok := s.Data[i].(string)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = sStringVal + strVal
			case operation == "GT":
				data[i] = sStringVal > strVal
			case operation == "GTE":
				data[i] = sStringVal >= strVal
			case operation == "LT":
				data[i] = sStringVal < strVal
			case operation == "LTE":
				data[i] = sStringVal <= strVal
			case operation == "EQ":
				data[i] = sStringVal == strVal
			case operation == "NEQ":
				data[i] = sStringVal != strVal
			}
		}

		if s.column.Dtype == base.DateTime {
			var tVal time.Time
			var duration time.Duration

			if operation != "ADD" {
				t, ok := curVal.(time.Time)
				if !ok {
					panic(errors.IncorrectTypedParameter("val", "time.Time"))
				}
				tVal = t
			} else {
				d, ok := curVal.(time.Duration)
				if !ok {
					panic(errors.IncorrectTypedParameter("val", "time.Duration"))
				}
				duration = d
			}
			sTVal, ok := s.Data[i].(time.Time)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			switch {
			case operation == "ADD":
				data[i] = sTVal.Add(duration)
			case operation == "GT":
				data[i] = sTVal.After(tVal)
			case operation == "LT":
				data[i] = sTVal.Before(tVal)
			case operation == "EQ":
				data[i] = sTVal.Equal(tVal)
			}
		}

		if s.column.Dtype == base.Bool {
			boolVal, ok := curVal.(bool)
			if !ok {
				panic(errors.IncorrectTypedParameter("val", "bool"))
			}
			sBoolVal, ok := s.Data[i].(bool)
			if !ok {
				panic(errors.InvalidSeriesValError(s.Data[i], i, s.column.Name))
			}
			nonPermittedOps := []interface{}{"ADD", "SUBTRACT", "GT", "LT", "GTE", "LTE"}
			switch {
			case helpers.Contains(nonPermittedOps, operation):
				panic(errors.SeriesDataTypeNotPermitted(operation, base.Bool))
			case operation == "AND":
				data[i] = boolVal && sBoolVal
			case operation == "OR":
				data[i] = boolVal || sBoolVal
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
	newS := helperCrud(s, val, "ADD", false)
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
	newS := helperCrud(s, val, "SUBTRACT", false)
	setOpFuncName(val, "subtract", s, newS)
	return newS
}

// Gt will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is greater than the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are greater than the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Gt(val interface{}) *Series {
	newS := helperCrud(s, val, "GT", true)
	setOpFuncName(val, "gt", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// Gte will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is greater than or equal to the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are greater than or equal to the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Gte(val interface{}) *Series {
	newS := helperCrud(s, val, "GTE", true)
	setOpFuncName(val, "gte", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// Lt will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is less than the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are lesser than the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Lt(val interface{}) *Series {
	newS := helperCrud(s, val, "LT", true)
	setOpFuncName(val, "lt", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// Lte will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is less than or equal to the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are lesser than or equal to the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Lte(val interface{}) *Series {
	newS := helperCrud(s, val, "LTE", true)
	setOpFuncName(val, "lte", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// Eq will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is equal to the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are equal to the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Eq(val interface{}) *Series {
	newS := helperCrud(s, val, "EQ", true)
	setOpFuncName(val, "eq", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// Neq will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports whether the Series values is not equal to the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be checked to see if they are not equal to the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) Neq(val interface{}) *Series {
	newS := helperCrud(s, val, "NEQ", true)
	setOpFuncName(val, "neq", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// AND will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports the value arising after Series value is anded against the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be anded against the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) AND(val interface{}) *Series {
	newS := helperCrud(s, val, "AND", true)
	setOpFuncName(val, "and", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// OR will compare the given value against each value in the calling Series and will return a new Series
// of type base.Bool, which reports the value arising after Series value is ored against the passed val.
//
// If another Series is passed as the val parameter using Col method, each value of the calling
// Series will be ored against the corresponding values in the passed Series.
//
// The function panics if incompatible values or an incompatible Series is passed.
func (s *Series) OR(val interface{}) *Series {
	newS := helperCrud(s, val, "OR", true)
	setOpFuncName(val, "or", s, newS)
	newS.column.Dtype = base.Bool
	return newS
}

// NOT will reverse each value in a base.Bool series to it's boolean counterpart.
//
// The function panics if the passed series is not of type base.Bool.
func (s *Series) NOT() *Series {
	newS := s.ShallowCopy()

	if s.column.Dtype != base.Bool {
		panic(errors.SeriesDataTypeNotPermitted("NOT", base.Bool))
	}

	data := make([]interface{}, s.Len())
	for i, val := range s.Data {
		boolVal, ok := val.(bool)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}
		data[i] = !boolVal
	}

	newS.Data = data
	newS.column.Name = helpers.FunctionNameWrapper("not", s.column.Name)
	return newS
}
