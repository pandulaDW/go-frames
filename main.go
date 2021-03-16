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

	df = df.Select("video_id", "trending_date", "title", "publish_time")

	capitalized := df.Col("title").Lower().Capitalized()
	df = df.WithColumnRenamed("title", capitalized)

	year := df.Col("publish_time").Year()
	df = df.WithColumnRenamed("year", year)

	fmt.Println(df.Head(4))
	fmt.Println(time.Since(start))
}
