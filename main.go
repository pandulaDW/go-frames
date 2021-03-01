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

	fmt.Println(df.Tail(5))
}
