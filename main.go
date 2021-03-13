package main

import (
	"errors"
	"fmt"
	"github.com/pandulaDW/go-frames/ioread"
	"log"
	"time"
)

func main() {
	df, err := ioread.ReadCSV(ioread.CsvOptions{
		Path:       "data/youtubevideos.csv",
		ParseDates: map[string][]string{time.RFC3339: {"publish_time"}, "06.02.01": {"trending_date"}},
	})

	if err != nil {
		log.Fatal(err)
	}

	truncate := func(val interface{}) (interface{}, error) {
		strVal, ok := val.(string)
		if !ok {
			return nil, errors.New("not a string")
		}
		if len(strVal) < 30 {
			return strVal, nil
		}
		return strVal[0:30] + "...", nil
	}

	avgViews := df.Data["views"].Avg()
	diffSeries, err := df.Data["views"].Apply(func(val interface{}) (interface{}, error) {
		intVal, ok := val.(int)
		if !ok {
			return nil, errors.New("not an int")
		}
		return (intVal - int(avgViews)) / int(avgViews), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	df.AddColumn(diffSeries)

	df, err = df.ApplyToColumns([]string{"title", "tags", "description"}, truncate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Head(4))
	fmt.Println(df.Data["trending_date"].Max())
}
