package series

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type opsTestSuite struct {
	suite.Suite
	SInt      *Series
	SFloat    *Series
	sBool     *Series
	SObject   *Series
	sDateTime *Series
}

func (suite *opsTestSuite) SetupTest() {
	suite.SInt = NewSeries("col", 443, 54, "", 90, 48, 82)
	suite.SFloat = NewSeries("col", 43.53, 21.1, 32.54, 65.75, nil)
	suite.SObject = NewSeries("col", "foo", "bar", "baz")
	suite.sDateTime = NewSeries("col", "2005-01-25", "", "2012-02-05", "1998-11-25", "2001-12-15")
	_ = suite.sDateTime.CastAsTime("2006-01-02")
}

func (suite *opsTestSuite) TestHelperCrud() {
	// assert that function panics for invalid series
	suite.PanicsWithError(errors.MismatchedNumOfRows(3, 2).Error(), func() {
		suite.SObject.Add(NewSeries("col", "foo", "bar"))
	})

	// assert that nil values will be skipped correctly
	suite.Nil(suite.SInt.Add(3).Data[2])
	suite.Equal(false, suite.SInt.Gt(3).Data[2])

	// INT ---------------------------------------
	// assert that function panics for incorrectly typed values for int
	suite.PanicsWithError(errors.IncorrectTypedParameter("val", "int").Error(), func() {
		suite.SInt.Add("foo")
	})

	// assert that function panics if invalid series value is encountered
	sInt := suite.SInt.DeepCopy()
	sInt.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		sInt.Add(3)
	})

	// FLOAT ---------------------------------------
	// assert that function panics for incorrectly typed values for float
	suite.PanicsWithError(errors.IncorrectTypedParameter("val", "float64").Error(), func() {
		suite.SFloat.Add(5)
	})

	// assert that function panics if invalid series value is encountered
	sFloat := suite.SFloat.DeepCopy()
	sFloat.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		sFloat.Add(3.2)
	})

	// Object ------------------------------------------
	// assert that function panics for incorrectly typed values for object
	suite.PanicsWithError(errors.IncorrectTypedParameter("val", "string").Error(), func() {
		suite.SObject.Add(23)
	})

	// assert that function panics if invalid series value is encountered
	sObject := suite.SObject.DeepCopy()
	sObject.Data[2] = 3.5
	suite.PanicsWithError(errors.InvalidSeriesValError(3.5, 2, "col").Error(), func() {
		sObject.Add("foo")
	})

	// DateTime ----------------------------------------
	// assert that function panics for incorrectly typed values for datetime
	suite.PanicsWithError(errors.IncorrectTypedParameter("val", "time.Duration").Error(), func() {
		suite.sDateTime.Add(23)
	})

	// assert that function panics if invalid series value is encountered
	sDateTime := suite.sDateTime.DeepCopy()
	sDateTime.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError(3.5, 2, "col").Error(), func() {
		sDateTime.Add(1 * time.Second)
	})
}

func (suite *opsTestSuite) TestSeries_Add() {
	// assert that function correctly returns an added series for int types
	suite.Equal(NewSeries("add(col, 4)", 447, 58, nil, 94, 52, 86), suite.SInt.Add(4))

	// assert that function correctly returns an added series when passed a series
	suite.Equal(NewSeries("add(col, test)", 445, 57, nil, 92, 52, 88), suite.SInt.Add(
		NewSeries("test", 2, 3, 1, 2, 4, 6)))

	// assert that function correctly returns an added series float types
	suite.Equal(NewSeries("add(col, 1.2)", 44.730000000000004, 22.3, 33.74, 66.95, nil),
		suite.SFloat.Add(1.2))

	// assert that function correctly returns an added series object types
	suite.Equal(NewSeries("add(col, -x)", "foo-x", "bar-x", "baz-x"), suite.SObject.Add("-x"))

	// assert that function correctly returns an added series datetime types
	expected := NewSeries("add(col, 24h0m0s)", "2005-01-26", "", "2012-02-06", "1998-11-26", "2001-12-16")
	_ = expected.CastAsTime("2006-01-02")
	suite.Equal(expected, suite.sDateTime.Add(24*time.Hour))
}

func TestOpsTestSuite(t *testing.T) {
	suite.Run(t, new(opsTestSuite))
}

// suite.sDateTime = NewSeries("col", "2005-01-25", "", "2012-02-05", "1998-11-25", "2001-12-15")
