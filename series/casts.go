package series

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	customErrors "github.com/pandulaDW/go-frames/errors"
	"strconv"
	"time"
)

// CastAsInt would take in a series and will return a series casted as base.Int.
//
// The function will cast a Series as below.
//
// base.Int - Return nil without modifying the Series
//
// base.Float - Round down the Float values to Integers.
//
// base.Bool - Convert true values to 1 and false values to 0.
//
// base.Object - Convert to integers by rounding down.
//
// base.DateTime - Returns an error
func (s *Series) CastAsInt() error {
	switch s.column.Dtype {
	case base.Int:
		return nil
	case base.DateTime:
		return customErrors.IncorrectDataType(base.DateTime)
	}

	for i, val := range s.Data {
		switch s.column.Dtype {
		case base.Float:
			floatVal, ok := val.(float64)
			if !ok {
				return customErrors.InvalidRowValue(i)
			}
			s.Data[i] = int(floatVal)
		case base.Object:
			intVal, err := strconv.ParseInt(fmt.Sprintf("%v", val), 10, 64)
			if err != nil {
				return customErrors.CustomWithStandardError(customErrors.InvalidRowValue(i).Error(), err)
			}
			s.Data[i] = intVal
		case base.Bool:
			boolVal, ok := val.(bool)
			if !ok {
				return customErrors.InvalidRowValue(i)
			}
			if boolVal {
				s.Data[i] = 1
			}
			s.Data[i] = 0
		}
	}
	return nil
}

// CastAsTime would iterate through each element in an object series
// and will check if all elements conform to the given layout format.
//
// The layout string has to be of time 2006-01-02T15:04:05Z (Jan 02). Any standard time constant
// format such as time.RFC850 and time.RFC822 can be passed as argument to the function
//
// function will return an error with the first invalid row number and if no errors are found
// the series will be registered as an base.DateTime type and the values will converted time.Time format
func (s *Series) CastAsTime(layout string) error {
	sData := make([]interface{}, 0, s.Len())

	if s.column.Dtype != base.Object {
		return errors.New("only a series with object type can be inferred as a datetime series")
	}

	for i, val := range s.Data {
		strValue, ok := val.(string)
		if !ok {
			return errors.New(fmt.Sprintf("value at row number %d is not a string", i))
		}

		t, err := time.Parse(layout, strValue)
		if err != nil {
			return errors.New(fmt.Sprintf("invalid value at line %d. %s", i+1, err.Error()))
		}

		sData = append(sData, t)
	}

	// if no errors are found set the data and dtype
	s.Data = sData
	s.column.Dtype = base.DateTime
	return nil
}
