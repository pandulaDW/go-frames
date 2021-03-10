package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type definitionTestSuite struct {
	suite.Suite
	S *Series
}

// Setting up the data for the test suite
func (suite *definitionTestSuite) SetupTest() {
	suite.S = NewSeries("col", 43.53, 21.1, 32.54, 65.75)
}

func (suite *definitionTestSuite) TestSeries_Len() {
	// assert that the length getter works correctly
	assert.Equal(suite.T(), 4, suite.S.Len())
}

func (suite *definitionTestSuite) TestSeries_GetColumn() {
	column := base.Column{Name: "col", Dtype: base.Float64}
	// assert that Column returned correctly
	assert.Equal(suite.T(), column, *suite.S.GetColumn())
}

func (suite *definitionTestSuite) TestSeries_SetColName() {
	copiedSeries := suite.S.DeepCopy()
	copiedSeries.SetColName("newColName")
	// assert that Column name set correctly
	assert.Equal(suite.T(), "newColName", copiedSeries.column.Name)
}

func (suite *definitionTestSuite) TestSeries_SetColIndex() {
	copiedSeries := suite.S.DeepCopy()
	copiedSeries.SetColIndex(3)
	// assert that Column index set correctly
	assert.Equal(suite.T(), 3, copiedSeries.column.ColIndex)
}

func TestDefinitionTestSuite(t *testing.T) {
	suite.Run(t, new(definitionTestSuite))
}
