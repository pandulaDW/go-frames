package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/constituents_csv.csv"})
	if err != nil {
		log.Fatal(err)
		return
	}

	m := df.Data["Sector"].ValueCounts()
	for key, val := range m {
		fmt.Printf("%v: %v\n", key, val)
	}
}
