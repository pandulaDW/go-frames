package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []interface{}{24, 54, 12, 90}

	// assert that the function returns -1 if the element is not found
	assert.Equal(t, -1, LinearSearch(11, arr))

	// assert that the function returns correct index if the element is found
	assert.Equal(t, 1, LinearSearch(54, arr))
}
