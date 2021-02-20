package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/ioread"
	"github.com/pandulaDW/go-frames/series"
)

func main() {
	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, 45.6, 23.12, 23.2)
	col4 := series.NewSeries("col4", "true", "false", "true", "true", "false")
	col5 := series.NewSeries("col5", 14, 12.23, 32.5, 64, 34.1)
	col6 := series.NewSeries("col6", 10, 12, "", 45, 89)

	_ = dataframes.NewDataFrame(col1, col2, col3, col4, col5, col6)
	df := ioread.ReadCSV(ioread.CsvOptions{Path: "data/iris.csv", Delimiter: ","})

	fmt.Println(df.Info())
}
