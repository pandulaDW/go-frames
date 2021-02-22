package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type columnsTestSuite struct {
	suite.Suite
	df               *DataFrame
	col1, col2, col3 *series.Series
}

// Setting up the data for the test suite
func (suite *columnsTestSuite) SetupTest() {
	suite.col1 = series.NewSeries("col1", 12, 34, 54, 65, 90)
	suite.col2 = series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	suite.col3 = series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.df = NewDataFrame(suite.col1, suite.col2, suite.col3)
}

func (suite *columnsTestSuite) TestDataFrame_Columns() {
	expected := []string{"col1", "col2", "col3"}
	// assert that column names are returned correctly
	suite.Equal(expected, suite.df.Columns())
}

func (suite *columnsTestSuite) TestDataFrame_SetColumnNames() {
	newColumnNames := []string{"newCol1", "newCol2", "newCol3"}
	copiedDF := suite.df.DeepCopy()
	copiedDF.SetColumnNames(newColumnNames)

	// assert that the column names are set correctly
	suite.Equal(copiedDF.Columns(), newColumnNames)

	// assert that the function will panic if mismatched number of column names are given
	suite.PanicsWithError("mismatched number of columns provided", func() {
		copiedDF.SetColumnNames(newColumnNames[0:2])
	})
}

func (suite *columnsTestSuite) TestDataFrame_RenameColumn() {
	// assert that function returns an error when column name is not found
	err := suite.df.RenameColumn("test", "newTest")
	suite.EqualError(err, "column name is not found")

	// assert that function correctly rename the column
	df := suite.df.DeepCopy()
	_ = df.RenameColumn("col3", "testCol")
	testCol := series.NewSeries("testCol", suite.col3.Data...)
	expected := NewDataFrame(suite.col1, suite.col2, testCol)
	suite.Equal(expected, df)
}

func TestColumnsTestSuite(t *testing.T) {
	suite.Run(t, new(columnsTestSuite))
}
