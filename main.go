package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{
		Path:       "data/youtubevideos.csv",
		ParseDates: map[string][]string{time.RFC3339: {"publish_time"}, "06.02.01": {"trending_date"}},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Info())
	fmt.Println(time.Since(start))
}
