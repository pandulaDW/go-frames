package ioread

import (
	"encoding/csv"
	"github.com/fatih/color"
	"github.com/pandulaDW/go-frames/dataframes"
	"io"
)

//CsvOptions describes the read options specific to only csv format
type CsvOptions struct {
	// Any valid string path is acceptable
	Path string
	// Delimiter to use. Default will be a ','. Must be a valid rune and must not be \r, \n
	Delimiter rune
	// Column to use as the index of the DataFrame. Default index will be used if unspecified
	IndexCol string
	// A list of columns, which should be casted as dates
	DateCols []string
	// The default date format to be used. This field is mandatory if the date cols field is specified
	DateFormat string
	// A map of date columns and their respected formats. This is useful if multiple date columns exists
	// with different formats.
	//
	// Format can be specified as the map key and list of column names can be given as map values.
	//
	// If both DateCols and ParseDates fields are present, DateCols field will be disregarded.
	ParseDates map[string][]string
	// If an encountered when parsing lines (e.g. a csv line with too many commas) will by default return an error
	// and no DataFrame will be returned. If true, then these “bad lines” will dropped from the DataFrame that is
	// returned.
	SkipErrorLines bool
	// WarnErrorLines prints a warning for each “error line” to the standard output. Default is set to false.
	WarnErrorLines bool
}

// injectCustomOptions will take in an csv options object and will return it with
// default options set if parameters were not provided.
func injectCustomOptions(options *CsvOptions) {
	if options.Delimiter == 0 {
		options.Delimiter = ','
	}
}

// ReadCSV reads a csv file using the given option arguments. Returns the created dataframe
// with an error, if any.
//
// Refer the CsvOptions struct to get more information on read arguments.
func ReadCSV(options CsvOptions) (*dataframes.DataFrame, error) {
	// inject the defaults
	injectCustomOptions(&options)

	// open the file
	file, err := fileHandling(options.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// helper variables
	var colNames []string
	isHeader := true
	content := make([][]string, 0)

	// create the reader
	reader := csv.NewReader(file)
	reader.Comma = options.Delimiter
	reader.LazyQuotes = true

	// reading the file
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			if options.SkipErrorLines {
				if options.WarnErrorLines {
					color.Red("%v", err)
				}
				continue
			}
			return nil, err
		}
		if isHeader {
			colNames = row
		} else {
			content = append(content, row)
		}
		isHeader = false
	}

	// convert the content to a dataframe
	df := dataframes.ConvertRowContentToDF(colNames, content)

	// set the index, if provided
	if df.ColumnExistsWithIndex(options.IndexCol) != -1 {
		df = df.SetIndex(options.IndexCol)
	}

	// parse the dates
	err = dateParsing(&options, df)
	if err != nil {
		return nil, err
	}

	return df, nil
}
