package helpers

import "time"

// IsTimeSet checks if time is actually set or has been added by Go
func IsTimeSet(t time.Time) bool {
	return t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0
}
