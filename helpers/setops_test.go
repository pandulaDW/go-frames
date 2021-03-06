package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifference(t *testing.T) {
	a := []interface{}{12, "foo", 64, 3.5, 53, 19}
	b := []interface{}{12, "foo", 34, 3.5, 53, "bar"}

	// assert that function returns the difference correctly
	assert.ElementsMatch(t, []interface{}{34, 64, 19, "bar"}, Difference(a, b))
}
