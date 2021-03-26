package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/nyc_air_bnb.csv", DateCols: []string{"last_review"},
		DateFormat: "2006-01-02"})

	if err != nil {
		log.Fatal(err)
	}

	df = df.FilterBySeries(df.Col("neighbourhood").Contains("Harlem"))
	fmt.Println(df.Head(5))
	fmt.Println("time took: ", time.Since(start))
}
