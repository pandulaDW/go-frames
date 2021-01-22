package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type testSuite struct {
	suite.Suite
	S *Series
}

// Setting up the data for the test suite
func (suite *testSuite) SetupTest() {
	suite.S = NewSeries("col", 43.53, 21.1, 32.54, 65.75)
}

func (suite *testSuite) TestSeries_Len() {
	assert.Equal(suite.T(), 4, suite.S.Len(), "Length getter should work correctly")
}

func (suite *testSuite) TestSeries_GetColumn() {
	column := base.Column{Name: "col", Dtype: base.Float}
	assert.Equal(suite.T(), column, *suite.S.GetColumn(), "Column returned correctly")
}

func (suite *testSuite) TestSeries_SetColName() {

}

func (suite *testSuite) TestSeries_SetColIndex() {
	column := base.Column{Name: "col", Dtype: base.Float}
	assert.Equal(suite.T(), column, *suite.S.GetColumn(), "Column returned correctly")
}

func TestDefinitionTestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
