package series

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/stretchr/testify/suite"
	"testing"
)

type createTestSuite struct {
	suite.Suite
}

func (suite *createTestSuite) TestSeries_InferType() {
	// assert that an integer type will be asserted correctly
	s := NewSeries("col", 12, 43, 54, 65, 76)
	suite.Equal(base.Int, s.column.Dtype, "dtype should be int")

	// assert that a float type will be asserted correctly
	s = NewSeries("col", 1.23, 4.46, 6.45, 7.34, 8.4)
	suite.Equal(base.Float, s.column.Dtype, "dtype should be float")

	// assert that an bool type will be asserted correctly
	s = NewSeries("col", true, false, false, true, false)
	suite.Equal(base.Bool, s.column.Dtype, "dtype should be boolean")

	// assert that an string type will be asserted correctly
	s = NewSeries("col", "foo", "bar", "miz")
	suite.Equal(base.Object, s.column.Dtype, "dtype should be object")

	// assert that a mixed type will be asserted as object
	s = NewSeries("col", 12, "foo", 54.21, "bar", true)
	suite.Equal(base.Object, s.column.Dtype, "dtype should be object")

	// assert that mix types of floats and int will be treated as float
	s = NewSeries("col", 12, 23.43, 54.32, 43.54, 23, 34.54, 5.6, 90)
	suite.Equal(base.Float, s.column.Dtype, "dtype should be float")
}

func (suite *createTestSuite) TestNewSeries() {
	// assert that a new series object will be created
	expected := NewSeries("newCol", 23, 43, 90, 87)
	expectedCol := base.Column{Name: "newCol", Dtype: base.Int, ColIndex: 0}
	actual := Series{column: expectedCol, Data: []interface{}{23, 43, 90, 87}}
	suite.Equal(*expected, actual, "NewSeries creates a Series object correctly")
}

func (suite *createTestSuite) TestSeries_Copy() {
	s := NewSeries("newCol", 23, 43, 90, 87)
	copied := s.Copy()

	// assert that two object references are different
	suite.NotEqual(fmt.Sprintf("%p", s), fmt.Sprintf("%p", copied),
		"two series are two different objects")

	// assert that the series objects are equal
	suite.Equal(*s, *copied, "series is copied correctly")
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}