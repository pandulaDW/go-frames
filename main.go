package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/series"
)

func main() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12)
	col4 := series.NewSeries("col4", true, false, true, true)

	df := dataframes.CreateDataFrame(col1, col2, col3, col4)
	fmt.Println(df)
	fmt.Println(df.Info())
	//df.Describe()
}
