package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/helpers"
	"time"
)

func main() {
	start := time.Now()
	//df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/supermarket_sales.csv", DateCols: []string{"Date"},
	//	DateFormat: "1/2/2006"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(df.Head(5))

	seq := helpers.GenerateRandomSeries(20, 10, 42, true)
	fmt.Println(seq)
	fmt.Println(len(seq))
	fmt.Println("time took: ", time.Since(start))
}
