package series

import "github.com/pandulaDW/go-frames/base"

//NewSeries will create a new series based on the column name and the
// variadic arguments given
func NewSeries(colName string, data ...interface{}) *Series {
	column := base.Column{Name: colName, Dtype: base.Object}
	seriesData := make([]interface{}, 0, len(data))

	for _, val := range data {
		seriesData = append(seriesData, val)
	}

	series := Series{column: column, Data: seriesData}

	// assert the type
	series.AssertType()

	return &series
}

// AssertType will assert the type of the series based on its data.
//
// If the column contains a mix of int types and float types, then that column will
// be considered as a float column.
func (s *Series) AssertType() {
	for _, val := range s.Data {
		switch val.(type) {
		case int:
			if s.column.Dtype == base.Float {
				continue
			}
			s.column.Dtype = base.Int
		case float64:
			s.column.Dtype = base.Float
		case bool:
			s.column.Dtype = base.Bool
		default:
			s.column.Dtype = base.Object
		}
	}
}
