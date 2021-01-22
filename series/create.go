package series

import "github.com/pandulaDW/go-frames/base"

//NewSeries will create a new series based on the column name and the
// variadic arguments given
func NewSeries(colName string, data ...interface{}) *Series {
	column := base.Column{Name: colName}
	seriesData := make([]interface{}, 0, len(data))

	for _, val := range data {
		seriesData = append(seriesData, val)
	}

	series := Series{column: column, Data: seriesData}

	// assert the type
	series.InferType()

	return &series
}

// InferType will assert the type of the series based on its data.
//
// If the column contains a mix of int types and float types, then that column will
// be considered as a float column.
func (s *Series) InferType() {
	for _, val := range s.Data {
		// if at least one value is object, the column will be set as object
		if s.column.Dtype == base.Object {
			break
		}
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

// Copy will create and return a NewSeries which is a deep copy of on the content of
// the current series
func (s *Series) Copy() *Series {
	newColumn := base.Column{Name: s.column.Name, Dtype: s.column.Dtype, ColIndex: s.column.ColIndex}
	newDataSlice := make([]interface{}, s.Len())

	// copy the data
	copy(newDataSlice, s.Data)

	// create and return the new series
	return &Series{column: newColumn, Data: newDataSlice}
}
