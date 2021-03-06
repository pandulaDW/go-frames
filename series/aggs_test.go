package series

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type aggregationTestSuite struct {
	suite.Suite
	SInt      *Series
	SFloat    *Series
	SMix      *Series
	SObject   *Series
	sDateTime *Series
}

// Setting up the data for the test suite
func (suite *aggregationTestSuite) SetupTest() {
	suite.SInt = NewSeries("col", 443, 54, "", 90, 48, 82)
	suite.SFloat = NewSeries("col", 43.53, 21.1, 32.54, 65.75, "")
	suite.SMix = NewSeries("col", 89, 69.1, 2.34, 1.58)
	suite.SObject = NewSeries("col", "foo", "bar")
	suite.sDateTime = NewSeries("col", "2005-01-25", "", "2012-02-05", "1998-11-25", "2001-12-15")
}

func (suite *aggregationTestSuite) TestSeries_Max() {
	// assert int series returns correct max value
	suite.Equal(443, suite.SInt.Max())

	// assert float series returns correct max value
	suite.Equal(65.75, suite.SFloat.Max())

	// assert mix series returns correct max value
	suite.Equal(float64(89), suite.SMix.Max())

	// assert datetime returns correct max value
	layout := "2006-01-02"
	_ = suite.sDateTime.CastAsTime(layout)
	t, _ := time.Parse(layout, "2012-02-05")
	suite.Equal(t, suite.sDateTime.Max())

	// assert that function panics when an invalid value is there for datetime dtype in index 0 and elsewhere
	suite.sDateTime.Data[0] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 0, "col").Error(), func() {
		suite.sDateTime.Max()
	})
	suite.sDateTime.Data[0] = time.Now()
	suite.sDateTime.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		suite.sDateTime.Max()
	})

	// assert that function panics when an invalid value is there for int dtype
	invalidIntSeries := suite.SInt.DeepCopy()
	invalidIntSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidIntSeries.Max()
	})

	// assert that function panics when an invalid value is there for float dtype
	invalidFloatSeries := suite.SFloat.DeepCopy()
	invalidFloatSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidFloatSeries.Max()
	})

	// assert returns correct nil interface as a default when dtype is not applicable
	suite.Nil(suite.SObject.Max())
}

func (suite *aggregationTestSuite) TestSeries_Min() {
	// assert int series returns correct min value
	suite.Equal(48, suite.SInt.Min())

	// assert float series returns correct min value
	suite.Equal(21.1, suite.SFloat.Min())

	// assert mix series returns correct min value
	suite.Equal(1.58, suite.SMix.Min())

	// assert datetime returns correct min value
	layout := "2006-01-02"
	_ = suite.sDateTime.CastAsTime(layout)
	t, _ := time.Parse(layout, "1998-11-25")
	suite.Equal(t, suite.sDateTime.Min())

	// assert that function panics when an invalid value is there for datetime dtype in index 0 and elsewhere
	suite.sDateTime.Data[0] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 0, "col").Error(), func() {
		suite.sDateTime.Min()
	})
	suite.sDateTime.Data[0] = time.Now()
	suite.sDateTime.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		suite.sDateTime.Min()
	})

	// assert that function panics when an invalid value is there for int dtype
	invalidIntSeries := suite.SInt.DeepCopy()
	invalidIntSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidIntSeries.Min()
	})

	// assert that function panics when an invalid value is there for float dtype
	invalidFloatSeries := suite.SFloat.DeepCopy()
	invalidFloatSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidFloatSeries.Min()
	})

	// assert returns correct nil interface as a default when dtype is not applicable
	suite.Nil(suite.SObject.Min())
}

func (suite *aggregationTestSuite) TestSeries_Sum() {
	// assert int series returns correct sum
	suite.Equal(float64(717), suite.SInt.Sum())

	// assert float series returns correct sum
	suite.Equal(162.92, suite.SFloat.Sum())

	// assert mix series returns correct sum
	suite.Equal(162.02, suite.SMix.Sum())

	// assert that function panics when the dtype is not applicable
	suite.PanicsWithError("sum can only be applied for a numerical series", func() {
		suite.SObject.Sum()
	})

	// assert that function panics when an invalid value is there for int dtype
	invalidIntSeries := suite.SInt.DeepCopy()
	invalidIntSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidIntSeries.Sum()
	})

	// assert that function panics when an invalid value is there for float dtype
	invalidFloatSeries := suite.SFloat.DeepCopy()
	invalidFloatSeries.Data[2] = "foo"
	suite.PanicsWithError(errors.InvalidSeriesValError("foo", 2, "col").Error(), func() {
		invalidFloatSeries.Sum()
	})
}

func (suite *aggregationTestSuite) TestSeries_Avg() {
	// assert int series returns correct avg
	suite.Equal(float64(717)/5, suite.SInt.Avg())

	// assert float series returns correct avg
	suite.Equal(162.92/4, suite.SFloat.Avg())

	// assert mix series returns correct avg
	suite.Equal(162.02/4, suite.SMix.Avg())
}

func TestAggregationTestSuite(t *testing.T) {
	suite.Run(t, new(aggregationTestSuite))
}
