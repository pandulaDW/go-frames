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
	Path      string
	Delimiter string
	IndexCol  string
}

// ReadCSV reads a csv file using the given option arguments. Returns the created dataframe
// with an error, if any.
//
// Refer the CsvOptions struct to get more information on read arguments.
func ReadCSV(options CsvOptions) (*dataframes.DataFrame, error) {
	file := fileHandling(options.Path)
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

	df := convertRowContentToDF(colNames, content)

	// set the index, if provided
	if df.IsColumnIncluded(options.IndexCol) != -1 {

	}

	return df, nil
}
