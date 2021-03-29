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

	df = df.WithColumn(df.Col("Total").Subtract(df.Col("Tax 5%")).Round(2))
	df = df.WithColumnRenamed("Payment", df.Col("Payment").Add(" System"))

	datetime := df.Col("Date").Add(" ").Add(df.Col("Time"))
	err = datetime.CastAsTime("1/2/2006 15:04")

	if err != nil {
		log.Fatal(err)
	}

	df = df.Drop("Date", "Time")

	cols := make([]string, 0)
	cols = append(cols, "Datetime")
	cols = append(cols, df.Columns()...)

	df = df.WithColumnRenamed("Datetime", datetime).ResetColumns(cols)

	fmt.Println(df.Head(5))
	fmt.Println("time took: ", time.Since(start))
}
