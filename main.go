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
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/nyc_air_bnb.csv", DateCols: []string{"last_review"},
		DateFormat: "2006-01-02"})

	if err != nil {
		log.Fatal(err)
	}

	s := series.NewSeries("col", "foo", nil, "bar", "", "baz")
	s.CountOfNA()

	df = df.RenameColumn("calculated_host_listings_count", "calc_host_count")
	fmt.Println(df.Head(4))

	fmt.Println("time took: ", time.Since(start))
}
