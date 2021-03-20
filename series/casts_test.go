package series

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	customErrors "github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSeries_CastAsInt(t *testing.T) {
	sData := NewSeries("col", "2016-02-06", "2017-12-26", "2015-05-05")
	_ = sData.CastAsTime("2006-01-02")

	// assert that correct response returns to incorrect series
	err := NewSeries("col", 12, 43, 64).CastAsInt()
	assert.Nil(t, err)
	err = sData.CastAsInt()
	assert.EqualError(t, err, customErrors.IncorrectDataType(base.DateTime).Error())

	// assert that float series will be correctly casted
	sFloat := NewSeries("col", 1.34, 54.12, 4.23, 5.6)
	sFloatCorrect := sFloat.DeepCopy()
	_ = sFloatCorrect.CastAsInt()
	assert.Equal(t, sFloatCorrect, NewSeries("col", 1, 54, 4, 5))

	// assert that error will be returned for a incorrect float
	sFloat.Data[2] = "foo"
	err = sFloat.CastAsInt()
	assert.NotNil(t, err)

	// assert that string series will be correctly casted
	sStr := NewSeries("col", "12", "43", "56", "21")
	sStrCorrect := sStr.DeepCopy()
	_ = sStrCorrect.CastAsInt()
	assert.Equal(t, sStrCorrect, NewSeries("col", 12, 43, 56, 21))

	// assert that error will be returned for a incorrect string
	sStr.Data[2] = "foo"
	sStr.column.Dtype = base.Object
	err = sStr.CastAsInt()
	assert.NotNil(t, err)

	// assert that string series will be correctly casted
	sBool := NewSeries("col", "True", "True", "False", "True")
	sBoolCorrect := sBool.DeepCopy()
	_ = sBoolCorrect.CastAsInt()
	assert.Equal(t, sBoolCorrect, NewSeries("col", 1, 1, 0, 1))

	// assert that error will be returned for a incorrect string
	sBool.Data[2] = "foo"
	err = sBool.CastAsInt()
	assert.NotNil(t, err)
}

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
