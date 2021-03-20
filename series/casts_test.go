package series

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//goland:noinspection GoNilness
func TestSeries_CastAsTime(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	sCorrect := NewSeries("col", "2016-02-06 13:11:24", "2017-12-26 03:11:22", "2015-05-05 16:17:32",
		"2016-02-06 06:23:24", "2018-01-03 12:10:24", "1996-01-06 08:12:22", "2011-07-08 12:14:45")
	sIncorrect := sCorrect.DeepCopy()
	sInt := NewSeries("col", 12, 34, 54, 66)

	// assert that the function will throw an error if non object type is given
	assert.Error(t, errors.New("only a series with object type can be inferred as a datetime series"),
		sInt.CastAsTime(layout))

	// assert that the function will throw an error if a value is not string
	sIncorrect.Data[2] = 12
	assert.Error(t, errors.New(fmt.Sprintf("value at row number %d is not a string", 2)),
		sIncorrect.CastAsTime(layout))

	// assert that the function will throw an error if there's a parsing issue
	sIncorrect.Data[2] = "2020:02:05"
	_, err := time.Parse(layout, "2020:02:05")
	assert.Error(t, errors.New(fmt.Sprintf("invalid value at row %d. %s", 2, err.Error())),
		sIncorrect.CastAsTime(layout))

	err = sCorrect.CastAsTime(layout)
	// assert that the function will return nil, if there aren't any issues
	assert.Nil(t, err)

	// assert that the dtype is set the base.DateTime
	assert.Equal(t, base.DateTime, sCorrect.column.Dtype)

	// assert that a row value can be inferred to time.Time correctly
	_, ok := sCorrect.Data[2].(time.Time)
	assert.Equal(t, true, ok)
}
