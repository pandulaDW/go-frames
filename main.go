package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/ioread"
	"github.com/pandulaDW/go-frames/series"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/iris.csv"})
	if err != nil {
		return
	}

	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	df = dataframes.NewDataFrame(col1, col2, col3, col4)
	df.SetIndex("col2")

	fmt.Println(df)
}
