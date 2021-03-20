package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/supermarket_sales.csv",
		DateCols: []string{"Date"}, DateFormat: "1/2/2006"})

	if err != nil {
		log.Fatal(err)
	}

	df = df.WithColumnRenamed("Month", df.Col("Date").Month())
	cols := make([]string, 2, len(df.Columns()))
	cols[0] = "Date"
	cols[1] = "Month"

	for _, col := range df.Columns() {
		if col != "Date" && col != "Month" {
			cols = append(cols, col)
		}
	}

	df = df.ResetColumns(cols)
	err = df.Col("Rating").CastAsInt()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Head(5))
	fmt.Println(time.Since(start))
}
