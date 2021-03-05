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
	suite.Equal(base.Int, s.column.Dtype)

	// assert that a float type will be asserted correctly
	s = NewSeries("col", 1.23, 4.46, 6.45, 7.34, 8.4)
	suite.Equal(base.Float, s.column.Dtype)

	// assert that an bool type will be asserted correctly
	s = NewSeries("col", true, false, false, true, false)
	suite.Equal(base.Bool, s.column.Dtype)

	// assert that an string type will be asserted correctly
	s = NewSeries("col", "foo", "bar", "miz")
	suite.Equal(base.Object, s.column.Dtype)

	// assert that a mixed type will be asserted as object
	s = NewSeries("col", 12, "foo", 54.21, "bar", true)
	suite.Equal(base.Object, s.column.Dtype)

	// assert that values such as slices and maps will be converted to strings
	s = NewSeries("col", 21, "bar", []int{12, 4}, map[string]int{"foo": 1})
	suite.Equal(base.Object, s.column.Dtype)

	// assert that mix types of floats and int will be treated as float
	s = NewSeries("col", 12, 23.43, 54.32, 43.54, 23, 34.54, 5.6, 90)
	suite.Equal(base.Float, s.column.Dtype)

	// assert that blanks in types other than string will not be treated as string
	s = NewSeries("col", "12", "23.43", "54.32", "", "43.54", "23", "34.54")
	suite.Equal(base.Float, s.column.Dtype)
}

func (suite *createTestSuite) TestConvertStringToTypedValue() {
	// assert that a string value will be returned as a string
	suite.Equal("foo", convertStringToTypedValue("foo"))

	// assert that an int value will be returned as an int
	suite.Equal(1224, convertStringToTypedValue("1224"))

	// assert that an float value will be returned as an float
	suite.Equal(12.24, convertStringToTypedValue("12.24"))

	// assert that an bool value will be returned as an bool
	suite.Equal(true, convertStringToTypedValue("true"))
}

func (suite *createTestSuite) TestNewSeries() {
	// assert that a new series object will be created
	expected := NewSeries("newCol", 23, 43, 90, 87)
	expectedCol := base.Column{Name: "newCol", Dtype: base.Int, ColIndex: 0}
	actual := Series{column: expectedCol, Data: []interface{}{23, 43, 90, 87}}
	suite.Equal(*expected, actual)
}

func (suite *createTestSuite) TestSeries_ShallowCopy() {
	s := NewSeries("newCol", 23, 43, 90, 87)
	copied := s.ShallowCopy()

	// assert that two object references are different
	suite.NotEqual(fmt.Sprintf("%p", s), fmt.Sprintf("%p", copied))

	// assert that two slice references are same
	suite.Equal(fmt.Sprintf("%p", s.Data), fmt.Sprintf("%p", copied.Data))

	// assert that the series objects are equal
	suite.Equal(*s, *copied)
}

func (suite *createTestSuite) TestSeries_DeepCopy() {
	s := NewSeries("newCol", 23, 43, 90, 87)
	copied := s.DeepCopy()

	// assert that two object references are different
	suite.NotEqual(fmt.Sprintf("%p", s), fmt.Sprintf("%p", copied))

	// assert that two slice references are different
	suite.NotEqual(fmt.Sprintf("%p", s.Data), fmt.Sprintf("%p", copied.Data))

	// assert that the series objects are equal
	suite.Equal(*s, *copied)
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(createTestSuite))
}
