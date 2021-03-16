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
	suite.PanicsWithError("mismatched number of columns provided. requires 2 columns, but 3 was provided",
		func() {
			copiedDF.SetColumnNames(newColumnNames[0:2])
		})
}

func (suite *columnsTestSuite) TestDataFrame_RenameColumn() {
	// assert that function panics when column name is not found
	suite.PanicsWithError("test column not found in the dataframe", func() {
		suite.df.RenameColumn("test", "newTest")
	})

	// assert that function correctly rename the column
	df := suite.df.DeepCopy().RenameColumn("col3", "testCol")
	testCol := series.NewSeries("testCol", suite.col3.Data...)
	expected := NewDataFrame(suite.col1, suite.col2, testCol)
	suite.Equal(expected, df)
}

func (suite *columnsTestSuite) TestDataFrame_ResetColumns() {
	// assert that function panics when incorrect number of column names are provided
	suite.PanicsWithError("mismatched number of columns provided. requires 2 columns, but 3 was provided",
		func() {
			suite.df.ResetColumns([]string{"col1", "col2"})
		})

	// assert that function panics when column name is not found
	suite.PanicsWithError("test column not found in the dataframe", func() {
		suite.df.ResetColumns([]string{"test", "col1", "col2"})
	})

	// assert that function correctly reorders the columns
	expected := NewDataFrame(suite.col2, suite.col3, suite.col1)
	suite.Equal(expected, suite.df.ShallowCopy().ResetColumns([]string{"col2", "col3", "col1"}))
}

func (suite *columnsTestSuite) TestDataFrame_Drop() {
	// assert that function panics when incorrect column names are provided
	suite.PanicsWithError("testCol column not found in the dataframe", func() {
		suite.df.ShallowCopy().Drop("col1", "testCol")
	})

	// assert that function correctly drops the columns
	expected := NewDataFrame(suite.col3)
	suite.Equal(expected, suite.df.Drop("col1", "col2"))
}

func (suite *columnsTestSuite) TestDataFrame_Select() {
	// assert that function panics when incorrect column names are provided
	suite.PanicsWithError("testCol column not found in the dataframe", func() {
		suite.df.Drop("col1", "testCol")
	})

	// assert that function correctly selects the columns
	expected := NewDataFrame(suite.col1, suite.col2)
	suite.Equal(expected, suite.df.Select("col1", "col2"))
}

func (suite *columnsTestSuite) TestDataFrame_ColumnExists() {
	// assert that function returns false when column does not exits
	suite.Equal(false, suite.df.ColumnExists("col4"))

	// assert that function returns true when column exits
	suite.Equal(true, suite.df.ColumnExists("col1"))
}

func (suite *columnsTestSuite) TestDataFrame_ColumnExistsWithIndex() {
	// assert that function returns -1 when column is empty
	suite.Equal(-1, suite.df.ColumnExistsWithIndex(""))

	// assert that function returns -1 when column is not found
	suite.Equal(-1, suite.df.ColumnExistsWithIndex("col5"))

	// assert that function returns correct index when column is found
	suite.Equal(1, suite.df.ColumnExistsWithIndex("col2"))
}

func TestColumnsTestSuite(t *testing.T) {
	suite.Run(t, new(columnsTestSuite))
}
