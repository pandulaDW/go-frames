package dataframes

import (
	"fmt"
	"unsafe"

	"github.com/pandulaDW/go-frames/helpers"
)

func (df *DataFrame) createInfoFooter() string {
	dtypes := make([]interface{}, 0, len(df.columns))

	for _, val := range df.columns {
		dtypes = append(dtypes, val.dtype)
	}

	valueCounts := helpers.ValueCounts(dtypes)
	dtypeStr := fmt.Sprintf("dtypes: float(%d), int(%d), object(%d), bool(%d)\n",
		valueCounts[Float], valueCounts[Int], valueCounts[Object], valueCounts[Bool])

	memSize := fmt.Sprintf("memory usage: %v bytes", unsafe.Sizeof(df))
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
