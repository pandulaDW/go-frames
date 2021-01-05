package main

import (
	"github.com/pandulaDW/go-frames/dataframes"
	orderedmap "github.com/wk8/go-ordered-map"
)

func main() {
	om := orderedmap.New()
	om.Set("col1", []int{12, 34, 54, 34})
	df := dataframes.CreateDataFrame(om)
	_ = df
}
