package main

import (
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"github.com/pandulaDW/go-frames/series"
	"log"
	"time"
)

func main() {
	start := time.Now()
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/supermarket_sales.csv",
		DateCols: []string{"Date"}, DateFormat: "1/2/2006"})

	if err != nil {
		log.Fatal(err)
	}

	df = df.WithColumn(df.Col("Date").Month())

	isFood := df.Col("Product line").Lower().Contains("food")
	df = df.FilterBySeries(isFood)

	middleClass, _ := df.Col("gross income").Apply(func(val interface{}) (interface{}, error) {
		if val.(float64) > 20 {
			return true, nil
		}
		return false, nil
	})

	df = df.FilterBySeries(middleClass)

	fmt.Println(df.Select("Payment", "gross income").Head(5))

	s := series.NewSeries("test", "2013-02-04 16:24:15", "2017-09-24 05:12:35", "2011-12-07 11:54:12")
	_ = s.CastAsTime("2006-01-02 15:04:05")

	t, _ := time.Parse("2006-01-02 15:04:05", "2013-02-04 16:24:15")
	fmt.Println(t.Second())

	fmt.Println("time took: ", time.Since(start))
}
