package dataframes

import "github.com/pandulaDW/go-frames/base"

// CreateDataFrame creates a dataframe using given data
func CreateDataFrame(data [][]interface{}, columns []string) *DataFrame {
	df := new(DataFrame)
	df.Data = make(DataFrameData)
	df.columns = make([]base.Column, 0)
	df.length = len(data[0])

	// Populate the dataframe and the columns
	for i, colData := range data {
		df.Data[columns[i]] = colData
		column := base.Column{Name: columns[i], ColIndex: i}
		df.columns = append(df.columns, column)
	}

	// infer the types
	df.assertType()

	return df
}
