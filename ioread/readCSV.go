package ioread

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// ReadCSV reads a csv file using the given option arguments.
func ReadCSV(options CsvOptions) {
	file := fileHandling(options.Path)

	// helper variables
	var colNames []string
	isHeader := false
	columnCount := 0
	content := make([][]string, 0)

	// reading the file
	scanner := bufio.NewScanner(file)
	rowNumber := 0
	for scanner.Scan() {
		row := scanner.Text()
		rowNumber++

		if isHeader {
			colNames = strings.Split(row, options.Delimiter)
			columnCount = len(colNames)
		} else {
			rowData := strings.Split(row, options.Delimiter)
			if len(rowData) != columnCount {
				panic(errors.New("mismatched number of columns in row number " + strconv.Itoa(rowNumber)))
			}
			content = append(content, rowData)
		}

		isHeader = true
	}

}
