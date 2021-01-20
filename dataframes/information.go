package dataframes

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pandulaDW/go-frames/base"

	"github.com/pandulaDW/go-frames/helpers"
)

func getRealSizeOf(v interface{}) int {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		panic(err)
	}
	return b.Len()
}

func (df *DataFrame) createInfoFooter() string {
	dtypes := make([]interface{}, 0, len(df.columns))

	for _, val := range df.columns {
		dtypes = append(dtypes, val.dtype)
	}

	valueCounts := helpers.ValueCounts(dtypes)
	dtypeStr := fmt.Sprintf("dtypes: float(%d), int(%d), object(%d), bool(%d)\n",
		valueCounts[base.Float], valueCounts[base.Int], valueCounts[base.Object], valueCounts[base.Bool])

	memSize := fmt.Sprintf("memory usage: %d bytes", getRealSizeOf(df.Data))
	return dtypeStr + memSize
}

// Info returns a dataframe containing information about the DataFrame including the
// index dtype and columns, non-null values and memory usage.
func (df *DataFrame) Info() string {
	columns := []string{"#", "Column", "Non-Null Count", "Dtype"}
	indices := make([]interface{}, 0)
	columnNames := make([]interface{}, 0)
	nonNulls := make([]interface{}, 0)
	dTypes := make([]interface{}, 0)

	for i := 0; i < len(df.columns); i++ {
		indices = append(indices, i+1)
		columnNames = append(columnNames, df.columns[i].name)
		nonNulls = append(nonNulls, fmt.Sprintf("%d non-null", df.length))
		dTypes = append(dTypes, df.columns[i].dtype)
	}

	data := make([][]interface{}, 4)
	data[0] = indices
	data[1] = columnNames
	data[2] = nonNulls
	data[3] = dTypes

	info := CreateDataFrame(data, columns)
	return info.String() + "\n" + df.createInfoFooter()
}

// Describe Generate descriptive statistics. including those that summarize the central tendency,
// dispersion and shape of a dataset’s distribution, excluding NA values. information would only be displayed
// for the numerical columns.
func (df *DataFrame) Describe() {
	columns := make([]Column, 0)

	// extract only the numerical columns
	for _, val := range df.columns {
		fmt.Println(val.dtype)
		if val.dtype == base.Int || val.dtype == base.Float {
			columns = append(columns, val)
		}
	}

	maxSeries := df.Agg(columns, helpers.MaxSeries)
	minSeries := df.Agg(columns, helpers.MinSeries)

	_ = maxSeries
	_ = minSeries

}
