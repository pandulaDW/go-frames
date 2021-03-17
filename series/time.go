package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
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
			data[i] = t.Month().String()
		case "DAY":
			data[i] = t.Day()
		}
	}

	return NewSeries("placeholder", data...)
}

// Year returns the year in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Year() *Series {
	year := helperTimeMethods(s, "YEAR")
	year.column.Dtype = base.Int
	year.column.Name = helpers.FunctionNameWrapper("year", s.column.Name)
	return year
}

// Month returns the month in which the value occurs in the Series in a string format.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Month() *Series {
	month := helperTimeMethods(s, "MONTH")
	month.column.Dtype = base.Object
	month.column.Name = helpers.FunctionNameWrapper("month", s.column.Name)
	return month
}

// Day returns the day of the month in which the value occurs in the Series.
//
// The function panics if the series type is not base.DateTime
func (s *Series) Day() *Series {
	day := helperTimeMethods(s, "DAY")
	day.column.Dtype = base.Int
	day.column.Name = helpers.FunctionNameWrapper("day", s.column.Name)
	return day
}
