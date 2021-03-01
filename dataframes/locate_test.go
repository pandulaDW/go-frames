package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type locateTestSuite struct {
	suite.Suite
	df, dfLocExpected *DataFrame
	col1, col2, col3  *series.Series
}

// Setting up the data for the test suite
func (suite *locateTestSuite) SetupTest() {
	suite.col1 = series.NewSeries("col1", 12, 34, 54, 65, 90)
	suite.col2 = series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	suite.col3 = series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.df = NewDataFrame(suite.col1, suite.col2, suite.col3)

	indices := []int{0, 1, 2}
	suite.dfLocExpected = NewDataFrame(suite.col1.Loc(indices), suite.col2.Loc(indices), suite.col3.Loc(indices))
}

func (suite *locateTestSuite) TestDataFrame_ILoc() {
	// assert that function panics when a wrong column name is given
	suite.PanicsWithError("testCol column is not found", func() {
		suite.df.ILoc([]int{1, 2, 3}, []string{"testCol"})
	})

	// assert that the function returns the correct dataframe
	suite.Equal(suite.dfLocExpected, suite.df.ILoc([]int{0, 1, 2}, suite.df.Columns()))
}

func (suite *locateTestSuite) TestDataFrame_Head() {
	// assert that function panics when n is higher
	suite.PanicsWithError("n cannot be higher than the length of the dataframe", func() {
		suite.df.Head(10)
	})

	// assert that the function returns the correct dataframe
	indices := []int{0, 1, 2}
	expected := NewDataFrame(suite.col1.Loc(indices), suite.col2.Loc(indices), suite.col3.Loc(indices))
	suite.Equal(expected, suite.df.Head(3))
}

func (suite *locateTestSuite) TestDataFrame_Tail() {
	// assert that function panics when n is higher
	suite.PanicsWithError("n cannot be higher than the length of the dataframe", func() {
		suite.df.Tail(10)
	})

	// assert that the function returns the correct dataframe
	indices := []int{2, 3, 4}
	expected := NewDataFrame(suite.col1.Loc(indices), suite.col2.Loc(indices), suite.col3.Loc(indices))
	expected.Index.Data = series.NewSeries("#", 2, 3, 4)
	suite.Equal(expected, suite.df.Tail(3))
}

func TestLocateTestSuite(t *testing.T) {
	suite.Run(t, new(locateTestSuite))
}
