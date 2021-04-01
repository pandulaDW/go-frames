package series

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/suite"
	"testing"
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

	// Object ---------------------------------------
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
}

func TestOpsTestSuite(t *testing.T) {
	suite.Run(t, new(opsTestSuite))
}
