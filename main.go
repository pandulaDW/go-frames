package main

import (
	"github.com/pandulaDW/go-frames/dataframes"
)

func main() {
	col1 := make([]interface{}, 0)
	col2 := make([]interface{}, 0)

	for _, v := range []int{12, 34, 54, 65} {
		col1 = append(col1, v)
	}

	for _, v := range []string{"foo", "bar", "raz"} {
		col2 = append(col2, v)
	}

	data := make([][]interface{}, 2)
	data[0] = col1
	data[1] = col2

	columns := []string{"col1", "col2"}
	df := dataframes.CreateDataFrame(data, columns)
	_ = df
}
