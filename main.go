package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/dataframes"
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

	fmt.Println(df.Head(2))
	fmt.Println(dataframes.ConvertMapToDataFrame(df.Col("Payment").ValueCounts()))

	fmt.Println("time took: ", time.Since(start))
}
