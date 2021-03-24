package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"strconv"
)

//NewSeries will create a new series based on the column name and the
// variadic arguments given
func NewSeries(colName string, data ...interface{}) *Series {
	column := base.Column{Name: colName}
	seriesData := make([]interface{}, 0, len(data))

	for _, val := range data {
		if fmt.Sprintf("%T", val) == "string" {
			val = convertStringToTypedValue(val.(string))
		}
		seriesData = append(seriesData, val)
	}

	series := Series{column: column, Data: seriesData}

	// assert the type
	series.InferType()

	return &series
}

// InferType will assert the type of the series based on its data.
//
// If the column contains a mix of int types and float types, then that column will
// be considered as a float column and the value will be converted to a float type.
//
// Date columns will be initiated as an object value and can be later cased as datetime.
//
// blank cells are considered as NA and if they are present in a numerical or boolean columns,
// column dtype will not be considered as an object.
func (s *Series) InferType() {
	for i, val := range s.Data {
		// if at least one value is object, the column will be set as object
		if s.column.Dtype == base.Object {
			s.Data[i] = fmt.Sprintf("%v", val)
			continue
		}
		switch val.(type) {
		case int:
			if s.column.Dtype == base.Float {
				s.Data[i] = float64(val.(int))
				continue
			}
			s.column.Dtype = base.Int
		case float64:
			// if int is already set, traverse back and cast every int to float
			if s.column.Dtype == base.Int {
				for j, intVal := range s.Data {
					value, ok := intVal.(int)
					if ok {
						s.Data[j] = float64(value)
					}
					if j == i-1 {
						break
					}
				}
			}
			s.column.Dtype = base.Float
		case bool:
			s.column.Dtype = base.Bool
		default:
			if val == "" && s.column.Dtype != base.Object {
				s.Data[i] = nil
				continue
			}
			s.column.Dtype = base.Object
			s.Data[i] = fmt.Sprintf("%v", val)
			// traverse back and change all previous values to string
			for j, otherVal := range s.Data {
				s.Data[j] = fmt.Sprintf("%v", otherVal)
				if j == i-1 {
					break
				}
			}
		}
	}
}

//convertStringToTypedValue will take in a string and will try to convert it
// to an int, float or a boolean. If it's a plain string, then the value will
// be left as it is
func convertStringToTypedValue(val string) interface{} {
	int64Val, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return int(int64Val)
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return floatVal
	}

	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		return boolVal
	}

	return val
}

// ShallowCopy will create and return a NewSeries which is a shallow copy of on the content of
// the current series. Underlying data source will be the same in both series.
func (s *Series) ShallowCopy() *Series {
	newColumn := base.Column{Name: s.column.Name, Dtype: s.column.Dtype, ColIndex: s.column.ColIndex}

	// create and return the new series
	return &Series{column: newColumn, Data: s.Data}
}

// DeepCopy will create and return a NewSeries which is a deep copy of on the content of
// the current series
func (s *Series) DeepCopy() *Series {
	newColumn := base.Column{Name: s.column.Name, Dtype: s.column.Dtype, ColIndex: s.column.ColIndex}
	newDataSlice := make([]interface{}, s.Len())

	// copy the data
	copy(newDataSlice, s.Data)

	// create and return the new series
	return &Series{column: newColumn, Data: newDataSlice}
}
