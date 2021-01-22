package series

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type aggregationTestSuite struct {
	suite.Suite
	SInt    *Series
	SFloat  *Series
	SMix    *Series
	SObject *Series
}

// Setting up the data for the test suite
func (suite *aggregationTestSuite) SetupTest() {
	suite.SInt = NewSeries("col", 443, 54, 90, 48, 82)
	suite.SFloat = NewSeries("col", 43.53, 21.1, 32.54, 65.75)
	suite.SMix = NewSeries("col", 89, 69.1, 2.34, 1.58)
	suite.SObject = NewSeries("col", "foo", "bar")
}

func (suite *aggregationTestSuite) TestSeries_Max() {
	// assert int series returns correct max value
	suite.Equal(443, suite.SInt.Max(), "int series returns correct max value")

	// assert float series returns correct max value
	suite.Equal(65.75, suite.SFloat.Max(), "float series returns correct max value")

	// assert mix series returns correct max value
	suite.Equal(float64(89), suite.SMix.Max(), "mix series returns correct max value")

	// assert returns correct nil interface as a default when dtype is not applicable
	suite.Nil(suite.SObject.Max(), "returns nil when the dtype is object")
}

func (suite *aggregationTestSuite) TestSeries_Min() {
	// assert int series returns correct min value
	suite.Equal(48, suite.SInt.Min(), "int series returns correct min value")

	// assert float series returns correct min value
	suite.Equal(21.1, suite.SFloat.Min(), "float series returns correct min value")

	// assert mix series returns correct min value
	suite.Equal(1.58, suite.SMix.Min(), "mix series returns correct min value")

	// assert returns correct nil interface as a default when dtype is not applicable
	suite.Nil(suite.SObject.Min(), "returns nil when the dtype is object")
}

func (suite *aggregationTestSuite) TestSeries_Sum() {
	// assert int series returns correct sum
	suite.Equal(float64(717), suite.SInt.Sum(), "int series returns correct sum")

	// assert float series returns correct sum
	suite.Equal(162.92, suite.SFloat.Sum(), "float series returns correct sum")

	// assert mix series returns correct sum
	suite.Equal(162.02, suite.SMix.Sum(), "mix series returns correct sum")

	// assert that function panics when the dtype is not applicable
	suite.PanicsWithError("sum can only be applied for a numerical series", func() {
		suite.SObject.Sum()
	}, "panics when the dtype is not applicable")
}

func (suite *aggregationTestSuite) TestSeries_Avg() {
	// assert int series returns correct avg
	suite.Equal(float64(717)/5, suite.SInt.Avg(), "int series returns correct avg")

	// assert float series returns correct avg
	suite.Equal(162.92/4, suite.SFloat.Avg(), "float series returns correct avg")

	// assert mix series returns correct avg
	suite.Equal(162.02/4, suite.SMix.Avg(), "mix series returns correct avg")
}

func TestAggregationTestSuite(t *testing.T) {
	suite.Run(t, new(aggregationTestSuite))
}
