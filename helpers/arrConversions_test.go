package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToIntArray(t *testing.T) {
	// assert that interface array is converted to an int array successfully
	data := []interface{}{1, 5, 10, 4}
	assert.Equal(t, []int{1, 5, 10, 4}, ToIntArray(data))
}

func TestToInterfaceFromInt(t *testing.T) {
	// assert that int array is converted to an interface array successfully
	data := []int{1, 5, 10, 4}
	assert.Equal(t, []interface{}{1, 5, 10, 4}, ToInterfaceFromInt(data))
}
