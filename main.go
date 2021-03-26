package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/nyc_air_bnb.csv", DateCols: []string{"last_review"},
		DateFormat: "2006-01-02"})

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`.* (\w+/\w+) .*`)
	reCheck := df.Col("name").RegexExtract(re, 1)

	df = df.WithColumnRenamed("extract", reCheck)
	df = df.FilterBySeries(reCheck.NotBlank())

	fmt.Println(df.Select("name", "extract").Head(5))
	fmt.Println("time took: ", time.Since(start))
}
