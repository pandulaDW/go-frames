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
	//df.data["title"] = df.data["title"].Lower().Capitalized()

	//year := df.data["publish_time"].Year()
	//year.SetColName("year")
	//
	//df = df.AddColumn(year)

	fmt.Println(df.Head(4))
	fmt.Println(time.Since(start))
}
