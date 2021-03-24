package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/nyc_air_bnb.csv"})

	if err != nil {
		log.Fatal(err)
	}

	df = df.RenameColumn("calculated_host_listings_count", "calc_host_count")

	fmt.Println(df.Head(5))
	fmt.Println(df.Info())

	x := []interface{}{2, 3, nil, "foo", 10}
	for _, el := range x {
		if el != nil {
			fmt.Println(el)
		}
	}

	fmt.Println("time took: ", time.Since(start))
}
