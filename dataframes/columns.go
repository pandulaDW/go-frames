package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
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

// SetColumnNames will rename the columns based on the column list provided in the given order.
// Panics if the length of the column name array is lower than the number of columns in the dataframe
func (df *DataFrame) SetColumnNames(cols []string) {
	if len(cols) != len(df.columns) {
		panic(errors.CustomError("mismatched number of columns provided"))
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

// RenameColumn will rename the column provided with the new column name.
//
// Return an error, if the col name is not found
func (df *DataFrame) RenameColumn(oldName, newName string) error {
	for _, col := range df.columns {
		if col.Name == oldName {
			s := df.Data[oldName]    // get underlying series
			col.Name = newName       // changing the column list
			delete(df.Data, oldName) // delete the dataframe map key
			s.SetColName(newName)    // changing the series colName
			df.Data[newName] = s     // set the new key
			return nil               // return nil
		}
	}

	return errors.CustomError("column name is not found")
}
