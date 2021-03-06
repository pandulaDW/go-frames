package dataframes

import (
	"github.com/pandulaDW/go-frames/base"
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
func (df *DataFrame) SetColumnNames(cols []string) *DataFrame {
	if len(cols) != len(df.columns) {
		panic(errors.MismatchedNumOfColumns(len(cols), len(df.columns)))
	}

	// creating a new dataframe and assigning it to the df
	newSeriesArray := make([]*series.Series, 0)
	for i, column := range df.columns {
		newSeries := df.Data[column.Name]
		newSeries.SetColName(cols[i])
		newSeriesArray = append(newSeriesArray, newSeries)
	}

	*df = *NewDataFrame(newSeriesArray...)
	return df
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

	panic(errors.ColumnNotFound(oldName))
}

// ResetColumns resets the column order of the dataframe. All the columns should be present
// and the function panics if a column is not found in the column list.
func (df *DataFrame) ResetColumns(columns []string) *DataFrame {
	currentColumns := helpers.ToInterfaceFromString(df.Columns())

	if len(columns) != len(currentColumns) {
		panic(errors.MismatchedNumOfColumns(len(columns), len(currentColumns)))
	}

	// modify the ColIndex of the columns
	for i, col := range columns {
		if index := helpers.LinearSearch(col, currentColumns); index != -1 {
			df.columns[index].ColIndex = i
		} else {
			panic(errors.ColumnNotFound(col))
		}
	}

	// modify the df.columns slice
	sort.SliceStable(df.columns, func(i, j int) bool {
		return df.columns[i].ColIndex < df.columns[j].ColIndex
	})

	return df
}

// Drop will drop the given columns from the dataframe. Column names can be specified
// as variadic arguments.
//
// Panics if any of the column name is not found
func (df *DataFrame) Drop(colNames ...string) *DataFrame {

	for _, colName := range colNames {
		if _, ok := df.Data[colName]; ok {
			delete(df.Data, colName)
		} else {
			panic(errors.ColumnNotFound(colName))
		}
	}

	filteredCols := make([]*base.Column, 0)
	for _, column := range df.columns {
		filteredCols = append(filteredCols, column)
	}

	df.columns = filteredCols
	return df
}

// IsColumnIncluded returns the index position of the column name provided.
// If the column is not found, the function will return -1.
func (df *DataFrame) IsColumnIncluded(colName string) int {
	for _, col := range df.columns {
		if col.Name == colName {
			return col.ColIndex
		}
	}
	return -1
}
