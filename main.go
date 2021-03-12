package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{
		Path:           "data/youtubevideos.csv",
		Delimiter:      ',',
		DateCols:       []string{"publish_time"},
		DateFormat:     time.RFC3339,
		SkipErrorLines: true,
		WarnErrorLines: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Info())
}

// TODO - add series display
