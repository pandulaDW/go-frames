package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
)

func main() {
	df, _ := ioread.ReadCSV(ioread.CsvOptions{Path: "data/A_data.csv"})
	_ = df.Data["date"].CastAsTime("2006-01-02")

	fmt.Println(df.Data["date"].Max())
	fmt.Println(df)
}

// TODO - check to see how to run series creation in different go-frames
// TODO - adjust display when number of rows are high
// TODO - add series display
