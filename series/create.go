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
	return &series
}
