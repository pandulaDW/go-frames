package series

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_GetMaxLength(t *testing.T) {
	s1 := NewSeries("col", "foo", "bar", "baz", "food")
	s2 := NewSeries("column", "foo", "bar", "baz", "food")
	s3 := NewSeries("col", "2010-05-01", "2015-11-21", "2010-03-01")
	_ = s3.CastAsTime("2006-01-02")

	// assert that column length will be returned based on the longest value
	assert.Equal(t, 4, s1.GetMaxLength())

	// assert that when column name is longer, it's length will be taken
	assert.Equal(t, 6, s2.GetMaxLength())

	// assert that date columns are properly handled
	assert.Equal(t, 11, s3.GetMaxLength())
}
