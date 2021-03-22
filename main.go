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

	df = df.WithColumn(df.Col("Date").Month())

	isFood := df.Col("Product line").Lower().Contains("food")
	df = df.FilterBySeries(isFood)

	t, _ := time.Parse("2006/01/02", "2019/01/01")
	fixedDiff, _ := df.Col("Date").DateDiff(t).Apply(func(val interface{}) (interface{}, error) {
		return val.(int) - 1, nil
	})

	df = df.WithColumnRenamed("diff", fixedDiff)

	fmt.Println(df.Head(3))

	fmt.Println("time took: ", time.Since(start))
}
