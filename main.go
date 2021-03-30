package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/ioread"
	"github.com/pandulaDW/go-frames/series"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/supermarket_sales.csv"})

	if err != nil {
		log.Fatal(err)
	}

	df = df.WithColumn(df.Col("Total").Subtract(df.Col("Tax 5%")).Round(2))
	df = df.WithColumnRenamed("Payment", df.Col("Payment").Add(" System"))

	datetime := df.Col("Date").Add(" ").Add(df.Col("Time"))
	err = datetime.CastAsTime("1/2/2006 15:04")

	if err != nil {
		log.Fatal(err)
	}

	df = df.WithColumnRenamed("Datetime", datetime).Drop("Date", "Time")
	df = df.MoveColumn("Datetime", 0)

	col1 := series.NewSeries("col1", 12, 34, 54, 65, 90)
	col2 := series.NewSeries("col2", "foo", "bar", "raz", "apple", "orange")
	col3 := series.NewSeries("col3", 54.31, 1.23, nil, 23.12, 23.2)
	col4 := series.NewSeries("col4", true, false, true, true, false)
	col5 := series.NewSeries("col5", "2013/04/05", "2023/03/01", "2013/01/05", "2009/07/15", "2011/02/01")
	col6 := series.NewSeries("col", "2010-05-01 15:21:20", "2010-05-01 10:02:22", "2010-05-01 13:25:22",
		"2010-05-01 02:41:12", "2012-03-41 23:21:02")
	_ = col5.CastAsTime("2006/01/02")
	_ = col6.CastAsTime("2006/01/02")
	df = dataframes.NewDataFrame(col1, col2, col3, col4, col5, col6)
	df = df.SetIndex("col2")

	fmt.Println(df)
	fmt.Println("time took: ", time.Since(start))
}
