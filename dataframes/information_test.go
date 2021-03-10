package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"testing"
)

type infoTestSuite struct {
	suite.Suite
	df *DataFrame
}

func (suite *infoTestSuite) SetupTest() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	col5 := series.NewSeries("col5", 14, 124, 32, 64, 34)
	col6 := series.NewSeries("col6", 24.31, 5.63, 78.3, 22.43, 43)
	col7 := series.NewSeries("col7", "2013-02-08", "2011-01-18", "2012-12-08", "2013-01-07", "2011-11-28")
	_ = col7.CastAsTime("2006-01-02")
	suite.df = NewDataFrame(col1, col2, col3, col4, col5, col6, col7)
}

func (suite *infoTestSuite) TestCreateInfoFooter() {
	var memSize int
	for _, col := range suite.df.Columns() {
		memSize += suite.df.Data[col].MemSize()
	}
	expected := fmt.Sprintf(
		"dtypes: float(2), int(2), object(1), datetime(1), bool(1)\nmemory usage: %d bytes", memSize)
	// assert that the info footer is created successfully
	suite.Equal(expected, suite.df.createInfoFooter())
}

func (suite *infoTestSuite) TestCreateInfoDF() {
	expected := NewDataFrame(
		series.NewSeries("Column", "col1", "col2", "col3", "col4", "col5", "col6", "col7"),
		series.NewSeries("Non-Null Count", helpers.RepeatIntoSlice("5 non-null", 7)...),
		series.NewSeries("Dtype",
			base.Int64, base.StringType, base.Float64, base.Boolean, base.Int64, base.Float64, base.DateTime))
	actual := suite.df.createInfoDF()
	// assert that the info body dataframe is created successfully
	suite.Equal(expected, actual)
}

func (suite *infoTestSuite) TestInfo() {
	expected := suite.df.createInfoDF().String() + "\n" + suite.df.createInfoFooter()
	// assert that the info body dataframe is created successfully
	suite.Equal(expected, suite.df.Info())
}

func (suite *infoTestSuite) TestDescribe() {
	expected := NewDataFrame(
		series.NewSeries("", "max", "min", "sum", "avg"),
		series.NewSeries("col1", 90, 12, float64(255), float64(51)),
		series.NewSeries("col3", 54.31, 1.23, 147.46, 29.49),
		series.NewSeries("col5", float64(124), float64(14), float64(268), 53.6),
		series.NewSeries("col6", 78.3, 5.63, 173.67, 34.73),
	)

	actual := suite.df.Describe()
	fmt.Println("data: ", expected.Columns(), actual.Columns())

	// assert that the describe body dataframe is created successfully
	suite.Equal(expected, actual)
}

func TestInfoTestSuite(t *testing.T) {
	suite.Run(t, new(infoTestSuite))
}
