package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/A_data.csv"})
	if err != nil {
		return
	}

	_ = df.Data["date"].CastAsTime("2006-01-02")
	fmt.Println(df.Data["date"].Min())

	fmt.Println(df.Tail(5))
}

// TODO - check to see how to run series creation in different go-frames
// TODO - add display tests
// TODO - adjust display when number of rows are high
// TODO - adjust display when number of rows are high
// TODO - add series display
