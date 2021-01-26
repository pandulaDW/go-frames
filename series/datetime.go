package series

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"time"
)

// CastAsTime would iterate through each element in an object series
// and will check if all elements conform to the given layout format.
//
// The layout string has to be of time 2006-02-01 15:04:05. Any standard time constant
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
			return errors.New(fmt.Sprintf("invalid value at row %d. %s", i, err.Error()))
		}

		sData = append(sData, t)
	}

	// if no errors are found set the data and dtype
	s.Data = sData
	s.column.Dtype = base.DateTime
	return nil
}
