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

// SetColumnNames will rename the columns based on the column list provided in the given order and
// will return a new dataframe.
//
// Panics if the length of the column name array is lower than the number of columns in the dataframe
func (df *DataFrame) SetColumnNames(cols []string) *DataFrame {
	if len(cols) != len(df.columns) {
		panic(errors.MismatchedNumOfColumns(len(cols), len(df.columns)))
	}

	// creating a new dataframe and assigning it to the df
	newSeriesArray := make([]*series.Series, 0)
	for i, column := range df.columns {
		newSeries := df.data[column.Name]
		newSeries = newSeries.SetColName(cols[i])
		newSeriesArray = append(newSeriesArray, newSeries)
	}

	modifiedDF := NewDataFrame(newSeriesArray...)
	return modifiedDF
}

// RenameColumn will rename the column provided with the new column name and returns
// a new dataframe.
//
// Panics if the col name is not found
func (df *DataFrame) RenameColumn(oldName, newName string) *DataFrame {
	copiedDF := df.ShallowCopy()

	for _, col := range copiedDF.columns {
		if col.Name == oldName {
			s := copiedDF.data[oldName]    // get underlying series
			col.Name = newName             // changing the column list
			delete(copiedDF.data, oldName) // delete the dataframe map key
			s.SetColName(newName)          // changing the series colName
			copiedDF.data[newName] = s     // set the new key
			return copiedDF                // return from the function
		}
	}

	panic(errors.ColumnNotFound(oldName))
}

// ResetColumns creates a new DataFrame and resets the column order of that DataFrame.
// All the columns should be present
// and the function panics if a column is not found in the column list.
func (df *DataFrame) ResetColumns(columns []string) *DataFrame {
	copiedDF := df.ShallowCopy()
	currentColumns := helpers.ToInterfaceFromString(copiedDF.Columns())

	if len(columns) != len(currentColumns) {
		panic(errors.MismatchedNumOfColumns(len(columns), len(currentColumns)))
	}

	// modify the ColIndex of the columns
	for i, col := range columns {
		if index := helpers.LinearSearch(col, currentColumns); index != -1 {
			copiedDF.columns[index].ColIndex = i
		} else {
			panic(errors.ColumnNotFound(col))
		}
	}

	// modify the copiedDF.columns slice
	sort.SliceStable(copiedDF.columns, func(i, j int) bool {
		return copiedDF.columns[i].ColIndex < copiedDF.columns[j].ColIndex
	})

	return copiedDF
}

// Drop will drop the given columns from the DataFrame and will return a new DataFrame.
// Column names can be specified as variadic arguments.
//
// Panics if any of the column name is not found
func (df *DataFrame) Drop(colNames ...string) *DataFrame {
	copiedDF := df.ShallowCopy()

	for _, colName := range colNames {
		if _, ok := copiedDF.data[colName]; ok {
			delete(copiedDF.data, colName)
		} else {
			panic(errors.ColumnNotFound(colName))
		}
	}

	diffCols := helpers.Difference(helpers.ToInterfaceFromString(colNames),
		helpers.ToInterfaceFromString(copiedDF.Columns()))

	filteredCols := make([]*base.Column, 0)
	for _, column := range copiedDF.columns {
		if helpers.LinearSearch(column.Name, diffCols) != -1 {
			filteredCols = append(filteredCols, column)
		}
	}

	// reset the indices
	for i, column := range filteredCols {
		column.ColIndex = i
		copiedDF.data[column.Name].GetColumn().ColIndex = i
	}

	copiedDF.columns = filteredCols
	return copiedDF
}

// Select returns a new DataFrame instance with the specified columns.
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

	modifiedDF := df.Drop(droppedCols...)
	return modifiedDF
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
	if _, ok := df.data[colName]; !ok {
		return -1
	}
	return df.data[colName].GetColumn().ColIndex
}

// MoveColumn is a helper method to move a column to the desired index position. It will place the column
// in the given index position and will shift the columns on the right side by one position.
//
// The function panics if the colName is not present or if index is out of bound
func (df *DataFrame) MoveColumn(colName string, index int) *DataFrame {
	idx := df.ColumnExistsWithIndex(colName)
	if idx == index {
		return df
	}
	if idx == -1 {
		panic(errors.ColumnNotFound(colName))
	}

	if index < 0 || index > len(df.columns) {
		panic(errors.CustomError("index is out of bound"))
	}

	// populate slice without the given column
	cols := make([]string, 0, len(df.columns)-1)
	for _, column := range df.columns {
		if column.Name != colName {
			cols = append(cols, column.Name)
		}
	}

	// creating a new slice with the resettled column
	resetCols := make([]string, len(df.columns))
	resetCols[index] = colName
	copy(resetCols, cols[:index])
	copy(resetCols[index+1:], cols[index:])

	// reset the dataframe and return
	newDf := df.ResetColumns(resetCols)
	return newDf
}
