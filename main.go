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

	df, err = df.ApplyToColumns([]string{"title", "tags", "description"}, truncate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df.Head(4))
}
