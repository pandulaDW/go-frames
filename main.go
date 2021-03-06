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
	col5 := series.NewSeries("col5", "2013/04/05", "2023/03/01", "2013/01/05", "2009/07/15", "2011/02/01")

	df := dataframes.NewDataFrame(col1, col2, col3, col4, col5)
	df.SetIndex("col2")
	_ = df.Data["col5"].CastAsTime("2006/01/02")
	fmt.Println(df)
}

// TODO - check to see how to run series creation in different go-frames
// TODO - add display tests
// TODO - adjust display when number of rows are high
// TODO - adjust display when number of rows are high
// TODO - add series display
