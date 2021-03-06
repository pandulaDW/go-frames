package dataframes

import (
	"github.com/pandulaDW/go-frames/errors"
	"github.com/pandulaDW/go-frames/series"
)

// WithColumn adds a new column to a new DataFrame object and returns that DataFrame.
//
// Panics if the column name of the series already exists or if
// length of the series is not equal to the length of the dataframe.
func (df *DataFrame) WithColumn(s *series.Series) *DataFrame {
	copiedDF := df.ShallowCopy()
	if s.Len() != copiedDF.Length() {
		panic(errors.MismatchedNumOfRows(copiedDF.length, s.Len()))
	}

	// set column index
	s = s.SetColIndex(len(copiedDF.columns))

	// append column
	copiedDF.columns = append(copiedDF.columns, s.GetColumn())

	// add the data
	copiedDF.data[s.GetColumn().Name] = s

	return copiedDF
}

// WithColumnRenamed either modifies an existing column with the provided Series or create a new Series
// depending on the column name. The function will return a new DataFrame object with the updated or added column.
//
// Panics if the length of the series is not equal to the length of the dataframe.
func (df *DataFrame) WithColumnRenamed(colName string, s *series.Series) *DataFrame {
	copiedDF := df.ShallowCopy()

	// rename the series column name
	s = s.SetColName(colName)

	// increase the column index and append column
	if _, ok := copiedDF.data[colName]; !ok {
		s = s.SetColIndex(len(copiedDF.columns))
		copiedDF.columns = append(copiedDF.columns, s.GetColumn())
	}

	if s.Len() != copiedDF.Length() {
		panic(errors.MismatchedNumOfRows(copiedDF.length, s.Len()))
	}

	// update the map
	copiedDF.data[colName] = s

	return copiedDF
}
