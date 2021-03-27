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

	janSales := df.Col("Date").Month().Contains("January")
	df = df.FilterBySeries(janSales)

	fmt.Println(df.Head(5))
	fmt.Println("time took: ", time.Since(start))
}
