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

	fmt.Println(df.Head(5))
	fmt.Println("time took: ", time.Since(start))
}
