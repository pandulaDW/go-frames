package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
	"sort"
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

// RenameColumn will rename the column provided with the new column name and returns
// the modified dataframe.
//
// Panics if the col name is not found
func (df *DataFrame) RenameColumn(oldName, newName string) *DataFrame {
	for _, col := range df.columns {
		if col.Name == oldName {
			s := df.Data[oldName]    // get underlying series
			col.Name = newName       // changing the column list
			delete(df.Data, oldName) // delete the dataframe map key
			s.SetColName(newName)    // changing the series colName
			df.Data[newName] = s     // set the new key
			return df                // return from the function
		}
	}

	panic(errors.CustomError("column name is not found"))
}

// ResetColumns resets the column order of the dataframe. All the columns should be present
// and the function panics if a column is not found in the column list.
func (df *DataFrame) ResetColumns(columns []string) *DataFrame {
	currentColumns := helpers.ToInterfaceFromString(df.Columns())

	// modify the ColIndex of the columns
	for i, col := range columns {
		if index := helpers.LinearSearch(col, currentColumns); index != -1 {
			df.columns[index].ColIndex = i
		} else {
			panic(errors.CustomError(col + " column is not found"))
		}
	}

	// modify the df.columns slice
	sort.SliceStable(df.columns, func(i, j int) bool {
		return df.columns[i].ColIndex < df.columns[j].ColIndex
	})

	return df
}

// TODO - complete above function
// TODO - add binary search
// TODO - change methods to return a dataframe
