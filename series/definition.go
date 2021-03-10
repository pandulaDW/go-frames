package series

import (
	"github.com/pandulaDW/go-frames/base"
)

type column interface {
	GetColumn() *base.Column
	SetColName(colName string)
	SetColIndex(colIndex int)
}

type Column struct {
	column base.Column
}

type Series interface {
	Len() int
}

type IntSeries struct {
	Data []base.Int
	Column
}

type FloatSeries struct {
	Data []base.Float
	Column
}

type BoolSeries struct {
	Data []base.Bool
	Column
}

type TimeSeries struct {
	Data []base.Time
	Column
}

type StringSeries struct {
	Data []base.String
	Column
}

// Len returns the length of the underlying series data
func (s *IntSeries) Len() int {
	return len(s.Data)
}

// Len returns the length of the underlying series data
func (s *FloatSeries) Len() int {
	return len(s.Data)
}

// Len returns the length of the underlying series data
func (s *BoolSeries) Len() int {
	return len(s.Data)
}

// Len returns the length of the underlying series data
func (s *TimeSeries) Len() int {
	return len(s.Data)
}

// Len returns the length of the underlying series data
func (s *StringSeries) Len() int {
	return len(s.Data)
}
