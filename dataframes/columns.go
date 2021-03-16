package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/helpers"
	"github.com/pandulaDW/go-frames/series"
	"sort"
)

// Col takes in a column name of the dataframe and returns the corresponding series.
//
// The function panics if the column name is not available in the column list
func (df *DataFrame) Col(colName string) *series.Series {
	s, ok := df.data[colName]
	if !ok {
		panic(errors.ColumnNotFound(colName))
	}
	return s
}

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
		newSeries := df.data[column.Name]
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
			s := df.data[oldName]    // get underlying series
			col.Name = newName       // changing the column list
			delete(df.data, oldName) // delete the dataframe map key
			s.SetColName(newName)    // changing the series colName
			df.data[newName] = s     // set the new key
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
		if _, ok := df.data[colName]; ok {
			delete(df.data, colName)
		} else {
			panic(errors.ColumnNotFound(colName))
		}
	}

	diffCols := helpers.Difference(helpers.ToInterfaceFromString(colNames),
		helpers.ToInterfaceFromString(df.Columns()))

	filteredCols := make([]*base.Column, 0)
	for _, column := range df.columns {
		if helpers.LinearSearch(column.Name, diffCols) != -1 {
			filteredCols = append(filteredCols, column)
		}
	}

	df.columns = filteredCols
	return df
}

// Select returns the dataframe instance with the specified columns.
// Column names can be specified as a variadic argument.
//
// The function panics if any of the given columns are not found in the dataframe.
func (df *DataFrame) Select(cols ...string) *DataFrame {
	diffCols := helpers.Difference(helpers.ToInterfaceFromString(cols),
		helpers.ToInterfaceFromString(df.Columns()))

	droppedCols := make([]string, 0, len(diffCols))
	for _, col := range diffCols {
		droppedCols = append(droppedCols, fmt.Sprintf("%v", col))
	}

	// will panic here, if the column is not found
	df.Drop(droppedCols...)
	return df
}

// ColumnExists returns true if the given column exists in the DataFrame. The function would
// return false otherwise.
func (df *DataFrame) ColumnExists(colName string) bool {
	if _, ok := df.data[colName]; !ok {
		return false
	}
	return true
}

// ColumnExistsWithIndex returns the index position of the column name provided.
// If the column is not found, the function will return -1.
func (df *DataFrame) ColumnExistsWithIndex(colName string) int {
	for _, col := range df.columns {
		if col.Name == colName {
			return col.ColIndex
		}
	}
	return -1
}
