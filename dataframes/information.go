package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"

	"github.com/pandulaDW/go-frames/helpers"
)

func (df *DataFrame) createInfoFooter() string {
	var memSize int
	dtypes := make([]interface{}, 0, len(df.columns))

	for _, col := range df.columns {
		dtypes = append(dtypes, col.Dtype)
		memSize += df.Data[col.Name].MemSize()
	}

	valueCounts := helpers.ValueCounts(dtypes)
	dtypeStr := fmt.Sprintf("dtypes: float(%d), int(%d), object(%d), bool(%d)\n",
		valueCounts[base.Float], valueCounts[base.Int], valueCounts[base.Object], valueCounts[base.Bool])

	memSizeStr := fmt.Sprintf("memory usage: %d bytes", memSize)
	return dtypeStr + memSizeStr
}

// Info returns a dataframe containing information about the DataFrame including the
// index dtype and columns, non-null values and memory usage.
func (df *DataFrame) Info() string {
	var indices, columnNames, nonNulls, dTypes []interface{}

	for i := 0; i < len(df.columns); i++ {
		indices = append(indices, i+1)
		columnNames = append(columnNames, df.columns[i].Name)
		nonNulls = append(nonNulls, fmt.Sprintf("%d non-null", df.length))
		dTypes = append(dTypes, df.columns[i].Dtype)
	}

	col0 := series.NewSeries("#", indices...)
	col1 := series.NewSeries("Column", columnNames...)
	col2 := series.NewSeries("Non-Null Count", nonNulls...)
	col3 := series.NewSeries("Dtype", dTypes...)

	info := NewDataFrame(col0, col1, col2, col3)
	return info.String() + "\n" + df.createInfoFooter()
}

// Describe Generate descriptive statistics. including those that summarize the central tendency,
// dispersion and shape of a dataset’s distribution, excluding NA values. information would only be displayed
// for the numerical columns.
func (df *DataFrame) Describe() {
	columns := make([]base.Column, 0)

	// extract only the numerical columns
	for _, val := range df.columns {
		if val.Dtype == base.Int || val.Dtype == base.Float {
			columns = append(columns, *val)
		}
	}

	maxSeries := df.Agg(columns, base.MAX)
	minSeries := df.Agg(columns, base.MIN)
	sumSeries := df.Agg(columns, base.SUM)
	avgSeries := df.Agg(columns, base.AVG)

	fmt.Println(maxSeries)
	fmt.Println(minSeries)
	fmt.Println(sumSeries)
	fmt.Println(avgSeries)
}

// TODO - Add non null columns properly
// TODO  - Refactor using series.loc later
