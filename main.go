package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
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

	df = df.FilterBySeries(notYangon)

	fmt.Println(df)

	fmt.Println("time took: ", time.Since(start))
}
