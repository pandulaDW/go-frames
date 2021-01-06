package main

import (
	"fmt"

	"github.com/pandulaDW/go-frames/dataframes"
)

func main() {
	col1 := make([]interface{}, 0)
	col2 := make([]interface{}, 0)
	col3 := make([]interface{}, 0)
	col4 := make([]interface{}, 0)

	for _, v := range []int{12, 34, 54, 65} {
		col1 = append(col1, v)
	}

	for _, v := range []string{"foo", "bar", "raz", "apple"} {
		col2 = append(col2, v)
	}

	for _, v := range []string{"foo2", "baz2", "oranges", "apple2"} {
		col3 = append(col3, v)
	}

	for _, v := range []bool{true, false, true, true} {
		col4 = append(col4, v)
	}

	data := make([][]interface{}, 4)
	data[0] = col1
	data[1] = col2
	data[2] = col3
	data[3] = col4

	columns := []string{"col1", "col2", "col3", "col4"}
	df := dataframes.CreateDataFrame(data, columns)

	fmt.Println(df)
}
