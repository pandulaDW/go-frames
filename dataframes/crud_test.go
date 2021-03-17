package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type crudTestSuite struct {
	suite.Suite
	df               *DataFrame
	col1, col2, col3 *series.Series
}

// Setting up the data for the test suite
func (suite *crudTestSuite) SetupTest() {
	suite.col1 = series.NewSeries("col1", 12, 34, 54, 65, 90)
	suite.col2 = series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	suite.col3 = series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.df = NewDataFrame(suite.col1, suite.col2, suite.col3)
}

func (suite *crudTestSuite) TestDataFrame_WithColumn() {
	// assert that the function panics when mismatched number of rows are added
	suite.PanicsWithError("mismatched number of rows provided. requires 5 rows, but 3 was provided", func() {
		suite.df.WithColumn(series.NewSeries("col3", true, false, true))
	})

	// assert that the function returns a new dataframe with the added column
	testCol := series.NewSeries("col4", true, false, true, false, false)
	expected := NewDataFrame(suite.col1, suite.col2, suite.col3, testCol)
	actual := suite.df.WithColumn(testCol)
	suite.Equal(expected, actual)
	suite.Equal(false, suite.df.IsEqual(actual))
}

func (suite *crudTestSuite) TestDataFrame_WithColumnRenamed() {
	// assert that the function panics when mismatched number of rows are added
	suite.PanicsWithError("mismatched number of rows provided. requires 5 rows, but 3 was provided", func() {
		suite.df.WithColumn(series.NewSeries("col3", true, false, true))
	})

	// assert that the function returns a new dataframe with the added column
	testCol := series.NewSeries("col4", true, false, true, false, false)
	renamedCol := testCol.SetColName("test")
	expected := NewDataFrame(suite.col1, suite.col2, suite.col3, renamedCol)
	actual := suite.df.WithColumnRenamed("test", testCol)
	fmt.Println(actual)
	fmt.Println(expected)
	suite.Equal(expected, actual)
	suite.Equal(false, suite.df.IsEqual(actual))
}

func TestCrudTestSuite(t *testing.T) {
	suite.Run(t, new(crudTestSuite))
}
