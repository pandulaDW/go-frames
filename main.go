package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	s := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/youtubevideos.csv",
		DateCols: []string{"publish_time"}, DateFormat: time.RFC3339})

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(df.Data["publish_time"].Max())
	fmt.Println(df.Info())
	fmt.Println(time.Since(s))
}

// TODO - add series display
