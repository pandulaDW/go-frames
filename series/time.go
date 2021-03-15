package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"time"
)

// helperTimeMethods returns a new Series based on the _type
func helperTimeMethods(s *Series, _type string) *Series {
	if s.column.Dtype != base.DateTime {
		panic(errors.IncorrectDataType(base.DateTime))
	}

	newS := s.ShallowCopy()

	for i, val := range s.Data {
		t, ok := val.(time.Time)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}
		switch _type {
		case "YEAR":
			newS.Data[i] = t.Year()
		case "MONTH":
			newS.Data[i] = t.Month()
		case "DAY":
			newS.Data[i] = t.Day()
		}
	}

	return newS
}

// Year returns the year in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Year() *Series {
	return helperTimeMethods(s, "YEAR")
}

// Month returns the month in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Month() *Series {
	return helperTimeMethods(s, "MONTH")
}

// Day returns the day of the month in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Day() *Series {
	return helperTimeMethods(s, "DAY")
}
