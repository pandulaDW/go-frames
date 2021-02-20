package ioread

import (
	"bufio"
	"github.com/pandulaDW/go-frames/dataframes"
	"github.com/pandulaDW/go-frames/errors"
	"io"
	"strconv"
	"strings"
)

// ReadCSV reads a csv file using the given option arguments.
func ReadCSV(options CsvOptions) *dataframes.DataFrame {
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
			panic(errors.CustomWithStandardError("error in reading line "+strconv.Itoa(rowNumber), err))
		}

		if isHeader {
			colNames = strings.Split(strings.TrimSpace(row), options.Delimiter)
			columnCount = len(colNames)
		} else {
			rowData := strings.Split(strings.TrimSpace(row), options.Delimiter)
			if len(rowData) != columnCount {
				panic(errors.CustomError("mismatched number of columns in trimmed number " + strconv.Itoa(rowNumber)))
			}
			content = append(content, rowData)
		}

		isHeader = false
	}

	return convertRowContentToDF(colNames, content)
}
