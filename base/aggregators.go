package base

// Aggregator defines the constant identifiers for the aggregator functions
type Aggregator string

const (
	// Max will return the maximum of a series
	MAX Aggregator = "MAX"

	// Min will return the minimum of a series
	MIN Aggregator = "MIN"

	// SUM will return the total of a series
	SUM Aggregator = "SUM"

	// AVG will return the average of a series
	AVG Aggregator = "AVG"
)
