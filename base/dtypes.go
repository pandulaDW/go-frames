package base

// DType contains the supported data type definitions
type DType string

const (
	// Object is Text or mixed numeric values
	Object DType = "Object"
	// Int is int typed numeric values
	Int DType = "Int"
	// Float is float32 typed floating point values
	Float DType = "Float"
	// Bool is True/False values
	Bool DType = "Bool"
	// DateTime is Date and Time values
	DateTime DType = "DateTime"
	// NA represents empty cells
	NA DType = "NA"
)
