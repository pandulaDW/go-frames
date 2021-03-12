package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/youtubevideos.csv",
		DateCols: []string{"publish_time"}, DateFormat: time.RFC3339})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Info())
}

// TODO - add series display
