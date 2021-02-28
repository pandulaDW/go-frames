package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type indexTestSuite struct {
	suite.Suite
	df               *DataFrame
	col1, col2, col3 *series.Series
}

// Setting up the data for the test suite
func (suite *indexTestSuite) SetupTest() {
	suite.col1 = series.NewSeries("col1", 12, 34, 54, 65, 90)
	suite.col2 = series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	suite.col3 = series.NewSeries("col3", 43.53, 21.1, 32.54, 65.75, 67.3)
	suite.df = NewDataFrame(suite.col1, suite.col2, suite.col3)
}

func (suite *indexTestSuite) TestDataFrame_SetIndex() {
	// assert that the function panics if column is not found
	suite.PanicsWithError("testCol column not found in the dataframe", func() {
		suite.df.SetIndex("testCol")
	})

	// assert that the function sets the index properly
	expected := NewDataFrame(suite.col2, suite.col3)
	expected.Index = Index{
		Data:     suite.col1,
		IsCustom: true,
	}
	actual := suite.df.ShallowCopy().SetIndex("col1")
	suite.Equal(expected, actual)
}

func (suite *indexTestSuite) TestDataFrame_SetIndexBySeries() {
	// assert that the function panics if the lengths are different
	s := series.NewSeries("col4", 42, 56, 12, 90)
	suite.PanicsWithError(errors.MismatchedNumOfRows(5, 4).Error(), func() {
		suite.df.SetIndexBySeries(s)
	})

	// assert that the function sets the index properly
	s = series.NewSeries("col4", 42, 56, 12, 90, 45)
	suite.Equal(Index{Data: s, IsCustom: true}, suite.df.ShallowCopy().SetIndexBySeries(s).Index)
}

func (suite *indexTestSuite) TestDataFrame_ResetIndex() {
	// assert that the function returns the dataframe if index is default
	suite.Equal(suite.df, suite.df.ResetIndex(true))

	// assert that the function adds a new column if drop is false
	df := suite.df.ShallowCopy().SetIndex("col3")
	suite.Equal(suite.df, df.ResetIndex(false))

	// assert that the function drops the index column if drop is true
	df = suite.df.ShallowCopy().SetIndex("col3")
	expected := suite.df.ShallowCopy().Drop("col3")
	actual := df.ResetIndex(true)
	suite.Equal(expected, actual)
}

func TestIndexTestSuite(t *testing.T) {
	suite.Run(t, new(indexTestSuite))
}
