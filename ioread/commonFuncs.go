package ioread

import (
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/errors"
	"os"
)

// fileHandling will handle the file and will return the file
func fileHandling(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.CustomWithStandardError("error in reading the file", err)
	}

	return file, nil
}

// dateParsing will parse the dates according to the options provided and returns an error if any.
func dateParsing(options *CsvOptions, df *dataframes.DataFrame) error {
	if options.DateCols == nil && options.ParseDates == nil {
		return nil
	}

	if options.DateCols != nil && options.ParseDates == nil {
		if options.DateFormat == "" {
			return errors.CustomError("DateFormat field should not be empty if DateCols field is present")
		}

		for _, col := range options.DateCols {
			if _, ok := df.Data[col]; !ok {
				return errors.ColumnNotFound(col)
			}
			err := df.Data[col].CastAsTime(options.DateFormat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
