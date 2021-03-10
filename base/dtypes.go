package base

import "time"

// DType contains the supported data type definitions
type DType string

const (
	// StringType is Text or mixed numeric values
	StringType DType = "StringType"
	// Int64 is int typed numeric values
	Int64 DType = "Int64"
	// Float64 is float64 typed floating point values
	Float64 DType = "Float64"
	// Boolean is True/False values
	Boolean DType = "Boolean"
	// DateTime is Date and Time values
	DateTime DType = "DateTime"
)

// Int is a box type for int64 type
type Int struct {
	Value int64
	IsNA  bool
}

// Float is a box type for float64 type
type Float struct {
	Value float64
	IsNA  bool
}

// Bool is a box type for bool type
type Bool struct {
	Value bool
	IsNA  bool
}

// Time is a box type for time.Time type
type Time struct {
	Value time.Time
	IsNA  bool
}

// String is a box type for string type
type String struct {
	Value string
	IsNA  bool
}
