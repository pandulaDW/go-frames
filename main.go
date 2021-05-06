package main

import (
	"fmt"
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

	dtSeries := df.Col("Date").Add(" ").Add(df.Col("Time"))
	err = dtSeries.CastAsTime("1/2/2006 15:04")
	if err != nil {
		log.Fatal(err)
	}

	df = df.Drop("Date", "Time")
	df = df.WithColumnRenamed("DateTime", dtSeries).MoveColumn("DateTime", 1)

	notYangon := df.Col("City").Eq("Yangon").NOT()

	t1, _ := time.Parse("2006-01-02", "2019-03-24")
	t2, _ := time.Parse("2006-01-02", "2019-03-28")

	isGreater := df.Col("DateTime").Gt(t1)
	isLower := df.Col("DateTime").Lt(t2)

	mask := series.ComposeWithAnd(notYangon, isGreater, isLower)

	df = df.FilterBySeries(mask)
	fmt.Println(df)

	fmt.Println("time took: ", time.Since(start))
}
