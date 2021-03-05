package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
)

type applyTestSuite struct {
	suite.Suite
	df               *DataFrame
	fun              ApplyFunc
	err              error
	col1, col2, col3 *series.Series
}

// Setting up the data for the test suite
func (suite *applyTestSuite) SetupTest() {
	suite.col1 = series.NewSeries("col1", 12, 34, 54, 65, 90)
	suite.col2 = series.NewSeries("col2", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.col3 = series.NewSeries("col3", 14, 124.23, 32, 64.65, 34)
	suite.df = NewDataFrame(suite.col1, suite.col2, suite.col3)
	suite.err = errors.New("not a float")
	suite.fun = func(val interface{}) (interface{}, error) {
		floatVal, ok := val.(float64)
		if !ok {
			return nil, suite.err
		}
		return math.Round(floatVal), nil
	}
}

func (suite *applyTestSuite) TestDataFrame_ApplyToRows() {
	// assert that function returns an error if encountered
	df, err := suite.df.ApplyToRows(suite.fun)
	suite.Nil(df)
	suite.EqualError(err, suite.err.Error())

	// assert that function returns a dataframe with nil if no error
	df = NewDataFrame(suite.col2)
	actual, err := df.ApplyToRows(suite.fun)
	suite.Nil(err)
	suite.NotNil(actual)
}

func TestApplyTestSuite(t *testing.T) {
	suite.Run(t, new(applyTestSuite))
}
