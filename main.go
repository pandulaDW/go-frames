package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/iris.csv", Delimiter: ","})
	if err != nil {
		return
	}

	df.SetIndex("species")
	df.ResetIndex(true)
	fmt.Println(df.Head(5))
}
