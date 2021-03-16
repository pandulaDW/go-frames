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

	data := make([]interface{}, s.Len())

	for i, val := range s.Data {
		t, ok := val.(time.Time)
		if !ok {
			panic(errors.InvalidSeriesValError(val, i, s.column.Name))
		}
		switch _type {
		case "YEAR":
			data[i] = t.Year()
		case "MONTH":
			data[i] = t.Month()
		case "DAY":
			data[i] = t.Day()
		}
	}

	return NewSeries("test", data...)
}

// Year returns the year in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Year() *Series {
	year := helperTimeMethods(s, "YEAR")
	year.column.Dtype = base.Int
	return year
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
