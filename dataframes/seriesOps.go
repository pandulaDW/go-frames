package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
)

//AggFuncType defines the function type for the aggregate function
type AggFuncType func(num []interface{}, dtype base.DType) interface{}

// Agg returns the maximum of the column specified. Panics if
// the column is not found
func (df *DataFrame) Agg(series []interface{}, dtype base.DType, aggregator AggFuncType) []interface{} {
	var aggSeries []interface{}
	aggSeries = append(aggSeries, aggregator(series, dtype))
	return aggSeries
}
