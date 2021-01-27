package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/series"
)

func main() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	col5 := series.NewSeries("col5", 14, 124, 32, 64, 34)
	col6 := series.NewSeries("col6", 24.31, 5.63, 78.3, 22.43, 43)
	col7 := series.NewSeries("col7", "2012-02-05", "2005-01-25", "1998-11-25", "2001-12-15", "2004-05-12")

	df := dataframes.NewDataFrame(col1, col2, col3, col4, col5, col6, col7)
	fmt.Println(df.Describe())

	var val interface{}
	val = "foo"
	assertedVal, ok := val.(float64)
	if !ok {
		assertedVal = float64(val.(int))
	}
	fmt.Println(assertedVal)
}
