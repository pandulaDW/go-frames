package dataframes

import (
	"github.com/pandulaDW/go-frames/series"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type displayTestSuite struct {
	suite.Suite
	df *DataFrame
}

// Setting up the data for the test suite
func (suite *displayTestSuite) SetupTest() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, nil, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	col5 := series.NewSeries("col5", "2013/04/05", "2023/03/01", "2013/01/05", "2009/07/15", "2011/02/01")
	col6 := series.NewSeries("col6", "2010-05-01 15:21:20", "2010-05-01 10:02:22", "2010-05-01 13:25:22",
		"2010-05-01 02:41:12", "2012-03-41 23:21:02")
	_ = col5.CastAsTime("2006/01/02")
	_ = col6.CastAsTime("2006/01/02")
	suite.df = NewDataFrame(col1, col2, col3, col4, col5, col6)
}

func (suite *displayTestSuite) TestString() {
	// assert that string representation matches
	expected := `
+-+----+------+-----+-----+-----------+-------------------+
| |col1|  col2| col3| col4|       col5|               col6|
+-+----+------+-----+-----+-----------+-------------------+
|0|  12|   foo|54.31| true| 2013-04-05|2010-05-01 15:21:20|
|1|  34|   bar| 1.23|false| 2023-03-01|2010-05-01 10:02:22|
|2|  54|   raz|  N/A| true| 2013-01-05|2010-05-01 13:25:22|
|3|  65| apple|23.12| true| 2009-07-15|2010-05-01 02:41:12|
|4|  90|orange| 23.2|false| 2011-02-01|2012-03-41 23:21:02|
+-+----+------+-----+-----+-----------+-------------------+
`
	expected = strings.TrimSpace(expected)
	suite.Equal(expected, suite.df.String())

	// assert that display string after changing index works as expected
	expected = `
+------+----+-----+-----+-----------+-------------------+
|      |col1| col3| col4|       col5|               col6|
|  col2|    |     |     |           |                   |
+------+----+-----+-----+-----------+-------------------+
|   foo|  12|54.31| true| 2013-04-05|2010-05-01 15:21:20|
|   bar|  34| 1.23|false| 2023-03-01|2010-05-01 10:02:22|
|   raz|  54|  N/A| true| 2013-01-05|2010-05-01 13:25:22|
| apple|  65|23.12| true| 2009-07-15|2010-05-01 02:41:12|
|orange|  90| 23.2|false| 2011-02-01|2012-03-41 23:21:02|
+------+----+-----+-----+-----------+-------------------+
`
	expected = strings.TrimSpace(expected)
	suite.Equal(expected, suite.df.SetIndex("col2").String())
}

func TestDisplayTestSuite(t *testing.T) {
	suite.Run(t, new(displayTestSuite))
}
