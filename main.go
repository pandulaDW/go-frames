package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{
		Path:       "data/A_data.csv",
		DateCols:   []string{"date"},
		DateFormat: "2006-01-02"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df)
}

// TODO - check to see how to run series creation in different go-frames
// TODO - add series display
