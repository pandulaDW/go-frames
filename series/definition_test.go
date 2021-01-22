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
	copiedSeries := suite.S.Copy()
	copiedSeries.SetColName("newColName")
	assert.Equal(suite.T(), "newColName", copiedSeries.column.Name, "Column name set correctly")
}

func (suite *testSuite) TestSeries_SetColIndex() {
	copiedSeries := suite.S.Copy()
	copiedSeries.SetColIndex(3)
	assert.Equal(suite.T(), 3, copiedSeries.column.ColIndex, "Column index set correctly")
}

func TestDefinitionTestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
