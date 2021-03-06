package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsTimeSet(t *testing.T) {
	dateOnly, _ := time.Parse("2006-01-02", "2015-06-05")
	dateAndTime := time.Now()

	// assert that the function returns true if only date is provided
	assert.Equal(t, true, IsTimeSet(dateOnly))

	// assert that the function returns false if both date and time is provided
	assert.Equal(t, false, IsTimeSet(dateAndTime))
}
