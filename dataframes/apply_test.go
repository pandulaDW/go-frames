package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/base"
	errors2 "github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
)

type applyTestSuite struct {
	suite.Suite
	df               *DataFrame
	fun1, fun2       base.ApplyFunc
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
	suite.fun1 = func(val interface{}) (interface{}, error) {
		floatVal, ok := val.(float64)
		if !ok {
			return nil, suite.err
		}
		return math.Round(floatVal), nil
	}
	suite.fun2 = func(val interface{}) (interface{}, error) {
		num, ok := val.(float64)
		if !ok {
			return nil, errors.New("not a float")
		}
		return math.Sqrt(num), nil
	}
}

func (suite *applyTestSuite) TestDataFrame_ApplyToColumns() {
	// assert that function returns an error if a column is not found
	df, err := suite.df.ApplyToColumns([]string{"col5"}, suite.fun2)
	suite.Nil(df)
	suite.EqualError(err, errors2.ColumnNotFound("col5").Error())

	// assert that function returns an error if found
	df, err = suite.df.ApplyToColumns([]string{"col1"}, suite.fun2)
	suite.Nil(df)
	suite.NotNil(err)

	// assert that function returns a dataframe with nil if no error
	df, err = suite.df.ApplyToColumns([]string{"col2"}, suite.fun2)
	suite.NotNil(df)
	suite.Nil(err)
}

func TestApplyTestSuite(t *testing.T) {
	suite.Run(t, new(applyTestSuite))
}
