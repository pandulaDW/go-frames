package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
)

func main() {
	df, _ := ioread.ReadCSV(ioread.CsvOptions{Path: "data/A_data.csv"})
	_ = df.Data["date"].CastAsTime("2006-01-02")

	fmt.Println(df.Select("open", "high"))
}

// TODO - check to see how to run series creation in different go-frames
// TODO - add series display
