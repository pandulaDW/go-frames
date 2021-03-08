package ioread

import (
	"bufio"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/errors"
	"io"
	"strconv"
	"strings"
)

//CsvOptions describes the read options specific to only csv format
type CsvOptions struct {
	// Any valid string path is acceptable
	Path string
	// Delimiter to use. Default will be a ","
	Delimiter string
	// Column to use as the index of the DataFrame. Default index will be used if unspecified
	IndexCol string
	// A list of columns, which should be casted as dates
	DateCols []string
	// The default date format to be used. This field is mandatory if the date cols field is specified
	DateFormatCommon string
	// A map of date columns and their respected formats. This is useful if multiple date columns exists
	// with different formats.
	//
	// Format can be specified as the map key and list of column names
	// can be given as map values.
	ParseDates map[string][]string
}

// injectCustomOptions will take in an csv options object and will return it with
// default options set if parameters were not provided.
func injectCustomOptions(options *CsvOptions) {
	if options.Delimiter == "" {
		options.Delimiter = ","
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
	columnCount := 0
	content := make([][]string, 0)

	// reading the file
	reader := bufio.NewReader(file)
	rowNumber := 0
	for {
		row, err := reader.ReadString('\n')
		rowNumber++
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.CustomWithStandardError("error in reading line "+strconv.Itoa(rowNumber), err)
		}

		if isHeader {
			colNames = strings.Split(strings.TrimSpace(row), options.Delimiter)
			columnCount = len(colNames)
		} else {
			rowData := strings.Split(strings.TrimSpace(row), options.Delimiter)
			if len(rowData) != columnCount {
				return nil, errors.CustomError("mismatched number of columns in row " + strconv.Itoa(rowNumber-1))
			}
			content = append(content, rowData)
		}
		isHeader = false
	}

	df := dataframes.ConvertRowContentToDF(colNames, content)

	// set the index, if provided
	if df.IsColumnIncluded(options.IndexCol) != -1 {
		df.SetIndex(options.IndexCol)
	}

	return df, nil
}
