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
		memSize += df.data[col.Name].MemSize()
	}

	valueCounts := helpers.ValueCounts(dtypes)
	dtypeStr := fmt.Sprintf("dtypes: float(%d), int(%d), object(%d), datetime(%d), bool(%d)\n",
		valueCounts[base.Float], valueCounts[base.Int], valueCounts[base.Object], valueCounts[base.DateTime],
		valueCounts[base.Bool])

	memSizeStr := fmt.Sprintf("memory usage: %s", helpers.ConvertSizeToString(memSize))
	return dtypeStr + memSizeStr
}

func (df *DataFrame) createInfoDF() *DataFrame {
	var indices, columnNames, nonNulls, dTypes []interface{}

	for i := 0; i < len(df.columns); i++ {
		indices = append(indices, i+1)
		columnNames = append(columnNames, df.columns[i].Name)
		nonNulls = append(nonNulls, fmt.Sprintf("%d non-null", df.length))
		dTypes = append(dTypes, df.columns[i].Dtype)
	}

	col1 := series.NewSeries("Column", columnNames...)
	col2 := series.NewSeries("Non-Null Count", nonNulls...)
	col3 := series.NewSeries("Dtype", dTypes...)

	info := NewDataFrame(col1, col2, col3)
	return info
}

// Info returns a dataframe containing information about the DataFrame including the
// index dtype and columns, non-null values and memory usage.
func (df *DataFrame) Info() string {
	info := df.createInfoDF()
	return info.String() + "\n" + df.createInfoFooter()
}

// Describe Generate descriptive statistics. including those that summarize the central tendency,
// dispersion and shape of a dataset’s distribution, excluding NA values. information would only be displayed
// for the numerical columns.
func (df *DataFrame) Describe() *DataFrame {
	colNames := make([]string, 0)

	// extract only the numerical columns
	for _, val := range df.columns {
		if val.Dtype == base.Int || val.Dtype == base.Float {
			colNames = append(colNames, val.Name)
		}
	}

	// create aggregation series
	maxSeries := series.NewSeries("max", df.Agg(colNames, base.MAX)...)
	minSeries := series.NewSeries("min", df.Agg(colNames, base.MIN)...)
	sumSeries := series.NewSeries("sum", df.Agg(colNames, base.SUM)...).Round(2, true)
	avgSeries := series.NewSeries("avg", df.Agg(colNames, base.AVG)...).Round(2, true)

	infoDF := NewDataFrame(maxSeries, minSeries, sumSeries, avgSeries)
	transposedInfo := infoDF.Transpose(true)

	// add initial column
	colNamesDescribe := make([]string, len(colNames)+1)
	colNamesDescribe[0] = ""
	copy(colNamesDescribe[1:], colNames)

	// set column names
	transposedInfo = transposedInfo.SetColumnNames(colNamesDescribe)

	return transposedInfo
}

// TODO - Add non null columns properly
// TODO  - Refactor using series.loc later
