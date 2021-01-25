package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/base"
)

// Agg returns the aggregated values of the columns specified with the given aggregator
// function identifier. Panics if the column or the identifier is not found
func (df *DataFrame) Agg(columns []string, aggregator base.Aggregator) []interface{} {
	var aggSeries []interface{}

	for _, col := range columns {
		s, ok := df.Data[col]
		if !ok {
			panic(errors.New(col + " not found in the dataframe"))
		}

		switch aggregator {
		case base.MAX:
			aggSeries = append(aggSeries, s.Max())
		case base.MIN:
			aggSeries = append(aggSeries, s.Min())
		case base.SUM:
			aggSeries = append(aggSeries, s.Sum())
		case base.AVG:
			aggSeries = append(aggSeries, s.Avg())
		default:
			panic(errors.New(string(aggregator + " is not a valid aggregator identifier")))
		}
	}

	return aggSeries
}
