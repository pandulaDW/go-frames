package series

import (
	"github.com/pandulaDW/go-frames/base"
)

type Series struct {
	data   []interface{}
	column base.Column
}

//NewSeries will create a new series based on the column name and the
// variadic arguments given
func NewSeries(colName string, data ...interface{}) *Series {
	column := base.Column{Name: colName, Dtype: base.Object}
	seriesData := make([]interface{}, 0, len(data))

	for _, val := range data {
		seriesData = append(seriesData, val)
	}

	series := Series{column: column, data: seriesData}
	return &series
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

// Data will return the underlying data of the series
func (s *Series) Data() []interface{} {
	return s.data
}
