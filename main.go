package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
)

func main() {
	df := ioread.ReadCSV(ioread.CsvOptions{Path: "data/iris.csv", Delimiter: ","})

	fmt.Println(df.Length())
	fmt.Println(df.Info())
	fmt.Println(df.Describe())
}
