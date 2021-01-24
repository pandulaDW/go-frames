package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type columnsTestSuite struct {
	suite.Suite
	df *DataFrame
}

// Setting up the data for the test suite
func (suite *columnsTestSuite) SetupTest() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.df = NewDataFrame(col1, col2, col3)
}

func (suite *columnsTestSuite) TestDataFrame_Columns() {
	expected := []string{"col1", "col2", "col3"}
	suite.Equal(expected, suite.df.Columns(), "column names are returned correctly")
}

func (suite *columnsTestSuite) TestDataFrame_SetColumnNames() {
	newColumnNames := []string{"newCol1", "newCol2", "newCol3"}
	copiedDF := suite.df.Copy()
	copiedDF.SetColumnNames(newColumnNames)

	// assert that the column names are set correctly
	suite.Equal(copiedDF.Columns(), newColumnNames, "column names are set correctly")

	// assert that the function will panic if mismatched number of column names are given
	suite.PanicsWithError("mismatched number of columns provided", func() {
		copiedDF.SetColumnNames(newColumnNames[0:2])
	}, "panics with mismatched error")
}

func TestColumnsTestSuite(t *testing.T) {
	suite.Run(t, new(columnsTestSuite))
}
