package base

// DType contains the supported data type definitions
type DType string

const (
	// String is Text or mixed numeric values
	String DType = "String"
	// Int is int typed numeric values
	Int DType = "Int"
	// Float is float64 typed floating point values
	Float DType = "Float"
	// Boolean is True/False values
	Boolean DType = "Boolean"
	// DateTime is Date and Time values
	DateTime DType = "DateTime"
)
