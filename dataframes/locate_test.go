package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type locateTestSuite struct {
	suite.Suite
	df, dfLocExpected *DataFrame
}

// Setting up the data for the test suite
func (suite *locateTestSuite) SetupTest() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	suite.df = NewDataFrame(col1, col2, col3)

	indices := []int{0, 1, 2}
	suite.dfLocExpected = NewDataFrame(col1.Loc(indices), col2.Loc(indices), col3.Loc(indices))
}

func (suite *locateTestSuite) TestDataFrame_Loc() {
	// assert that function panics when a wrong column name is given
	suite.PanicsWithError("testCol column is not found", func() {
		suite.df.Loc([]int{1, 2, 3}, []string{"testCol"})
	})

	// assert that the function returns the correct dataframe
	suite.Equal(suite.dfLocExpected, suite.df.Loc([]int{0, 1, 2}, suite.df.Columns()))
}

func TestLocateTestSuite(t *testing.T) {
	suite.Run(t, new(locateTestSuite))
}
