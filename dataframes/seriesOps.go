package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
)

//AggFuncType defines the function type for the aggregate function
type AggFuncType func(num []interface{}, dtype base.DType) interface{}

// Agg returns the maximum of the column specified. Panics if
// the column is not found
func (df *DataFrame) Agg(columns []base.Column, aggregator AggFuncType) []interface{} {
	var aggSeries []interface{}

	for _, col := range columns {
		aggSeries = append(aggSeries, aggregator(df.Data[col.Name], col.Dtype))
	}

	return aggSeries
}
