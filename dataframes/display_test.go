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
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	suite.df = NewDataFrame(col1, col2, col3, col4)
}

func (suite *displayTestSuite) TestString() {
	expected := `
+-+----+------+-----+-----+
| |col1|  col2| col3| col4|
+-+----+------+-----+-----+
|0|  12|   foo|54.31| true|
|1|  34|   bar| 1.23|false|
|2|  54|   raz| 45.6| true|
|3|  65| apple|23.12| true|
|4|  90|orange| 23.2|false|
+-+----+------+-----+-----+
`
	expected = strings.TrimSpace(expected)
	// assert that string representation matches
	suite.Equal(expected, suite.df.String())

	// assert that the index is correctly represented when displaying
	expected = `
+------+----+-----+-----+
|      |col1| col3| col4|
|  col2|    |     |     |
+------+----+-----+-----+
|   foo|  12|54.31| true|
|   bar|  34| 1.23|false|
|   raz|  54| 45.6| true|
| apple|  65|23.12| true|
|orange|  90| 23.2|false|
+------+----+-----+-----+
`
	expected = strings.TrimSpace(expected)
	suite.Equal(expected, suite.df.ShallowCopy().SetIndex("col2").String())
}

func TestDisplayTestSuite(t *testing.T) {
	suite.Run(t, new(displayTestSuite))
}
