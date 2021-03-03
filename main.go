package main

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"math"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{Path: "data/iris.csv"})
	if err != nil {
		return
	}

	//df.Drop("species")
	roundTo2 := func(val interface{}) (interface{}, error) {
		floatVal, ok := val.(float64)
		if !ok {
			return nil, errors.New("not a good value")
		}
		return math.Round(floatVal), nil
	}

	df, err = df.ApplyToRows(roundTo2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(df.Head(5))
	}
}
