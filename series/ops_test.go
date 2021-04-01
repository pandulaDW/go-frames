package series

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type opsTestSuite struct {
	suite.Suite
	SInt      *Series
	SFloat    *Series
	sBool     *Series
	SObject   *Series
	sDateTime *Series
}

func (suite *opsTestSuite) SetupTest() {
	suite.SInt = NewSeries("col", 443, 54, "", 90, 48, 82)
	suite.SFloat = NewSeries("col", 43.53, 21.1, 32.54, 65.75, nil)
	suite.SObject = NewSeries("col", "foo", "bar", "baz")
	suite.sDateTime = NewSeries("col", "2005-01-25", "", "2012-02-05", "1998-11-25", "2001-12-15")
}

func (suite *opsTestSuite) TestHelperCrud() {
	// assert that invalid series will be panicked
	suite.PanicsWithError(errors.MismatchedNumOfRows(3, 2).Error(), func() {
		suite.SObject.Add(NewSeries("col", "foo", "bar"))
	})

	// assert that nil values will be skipped correctly
	suite.Nil(suite.SInt.Add(3).Data[2])
	suite.Equal(false, suite.SInt.Gt(3).Data[2])
}

func TestOpsTestSuite(t *testing.T) {
	suite.Run(t, new(opsTestSuite))
}
