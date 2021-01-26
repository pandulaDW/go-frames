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
	col7 := series.NewSeries("col7", 25, 56, 12, 27, 59)

	df := dataframes.NewDataFrame(col1, col2, col3, col4, col5, col6, col7)
	fmt.Println(df.Info())
}
