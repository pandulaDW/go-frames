package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
)

// Agg returns the aggregated values of the columns specified with the given aggregator
// function identifier. Panics if the column or the identifier is not found
func (df *DataFrame) Agg(columns []base.Column, aggregator base.Aggregator) []interface{} {
	var aggSeries []interface{}

	for _, col := range columns {
		s, ok := df.Data[col.Name]
		if !ok {
			panic(col.Name + " not found in the dataframe")
		}

		switch aggregator {
		case base.MAX:
			aggSeries = append(aggSeries, s.Max())
		case base.MIN:
			aggSeries = append(aggSeries, s.Min())
		default:
			panic(aggregator + " is not a valid aggregator identifier")
		}
	}

	return aggSeries
}
