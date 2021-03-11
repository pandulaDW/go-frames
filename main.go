package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"time"
)

func main() {
	s := time.Now()
	//df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/youtubevideos.csv",
	//	DateCols: []string{"publish_time"}, DateFormat: time.RFC3339})

	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/irisIncorrect.csv",
		SkipErrorLines: true, WarnErrorLines: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(df)
	fmt.Println(time.Since(s))
}

// TODO - check to see how to run series creation in different go-frames
// TODO - add series display
