package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
)

// AddColumn adds a new column to a new DataFrame object and returns that DataFrame.
//
// Panics if the length of the series is not equal to the length of the dataframe.
func (df *DataFrame) AddColumn(s *series.Series) *DataFrame {
	if s.Len() != df.Length() {
		panic(errors.MismatchedNumOfRows(df.length, s.Len()))
	}

	// append column
	df.columns = append(df.columns, s.GetColumn())

	// add the data
	df.data[s.GetColumn().Name] = s

	return df
}

// WithColumnRenamed either modifies an existing column with the provided Series or create a new Series
// depending on the column name. The function will return a new DataFrame object with the updated or added column.
//
// Panics if the length of the series is not equal to the length of the dataframe.
func (df *DataFrame) WithColumnRenamed(colName string, s *series.Series) *DataFrame {
	copiedDF := df.ShallowCopy()

	// rename the series column name
	s.SetColName(colName)

	// increase the column index and append column
	if _, ok := copiedDF.data[colName]; !ok {
		s.SetColIndex(len(copiedDF.columns) + 1)
		copiedDF.columns = append(copiedDF.columns, s.GetColumn())
	}

	if s.Len() != copiedDF.Length() {
		panic(errors.MismatchedNumOfRows(copiedDF.length, s.Len()))
	}

	// update the map
	copiedDF.data[colName] = s

	return copiedDF
}
