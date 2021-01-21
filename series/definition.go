package series

import (
	"github.com/pandulaDW/go-frames/base"
)

type Series struct {
	Data   []interface{}
	column base.Column
}

//Len returns the length of the underlying series data
func (s *Series) Len() int {
	return len(s.Data)
}

// GetColumn will return the column type of the series
func (s *Series) GetColumn() *base.Column {
	return &s.column
}

//SetColName will set the column name of the series
func (s *Series) SetColName(colName string) {
	s.column.Name = colName
}

//SetColIndex will set the column index of the series
func (s *Series) SetColIndex(colIndex int) {
	s.column.ColIndex = colIndex
}
