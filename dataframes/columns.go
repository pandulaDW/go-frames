package dataframes

import (
	"errors"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/series"
)

// Columns returns the column names of the dataframe
func (df *DataFrame) Columns() []string {
	names := make([]string, len(df.columns))
	for i, val := range df.columns {
		names[i] = val.Name
	}
	return names
}

// ColDType returns the datatype of the column. If the column is not
// found, it will return an error with an empty string
func (df *DataFrame) ColDType(colName string) (base.DType, error) {
	for _, val := range df.columns {
		if val.Name == colName {
			return val.Dtype, nil
		}
	}
	return "", errors.New("column not found")
}

// SetColumnNames will rename the columns based on the column list provided in the given order.
// Panics if the length of the column name array is lower than the number of columns in the dataframe
func (df *DataFrame) SetColumnNames(cols []string) {
	if len(cols) != len(df.columns) {
		panic(errors.New("mismatched number of columns provided"))
	}

	// creating a new dataframe and assigning it to the df
	newSeriesArray := make([]*series.Series, 0)
	for i, column := range df.columns {
		newSeries := df.Data[column.Name]
		newSeries.SetColName(cols[i])
		newSeriesArray = append(newSeriesArray, newSeries)
	}

	*df = *NewDataFrame(newSeriesArray...)
}
