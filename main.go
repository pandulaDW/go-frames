package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/supermarket_sales.csv", DateCols: []string{"Date"},
		DateFormat: "1/2/2006"})
	if err != nil {
		log.Fatal(err)
	}

	df = df.Sample(10, 40, false)
	fmt.Println(df)

	fmt.Println("time took: ", time.Since(start))
}
