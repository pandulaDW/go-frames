package dataframes

// CreateDataFrame creates a dataframe using given data
func CreateDataFrame(data [][]interface{}, columns []string) *DataFrame {
	df := new(DataFrame)
	df.Data = make(DataFrameData)
	df.columns = make([]Column, 0)
	df.length = len(data[0])

	// Populate the dataframe and the columns
	for i, colData := range data {
		df.Data[columns[i]] = colData
		column := Column{name: columns[i], colIndex: i}
		df.columns = append(df.columns, column)
	}

	// infer the types
	df.assertType()

	return df
}
