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

	loc, _ := time.LoadLocation("Australia/Canberra")
	fmt.Println("time now: ", time.Now())
	fmt.Println("time in Canberra: ", time.Now().In(loc))
	fmt.Println("time difference")

	fmt.Println("time took: ", time.Since(start))
}
