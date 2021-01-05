package dataframes

import (
	"fmt"
	"strings"

	"github.com/pandulaDW/go-frames/helpers"
)

// returns the max column size by considering the whole column and the column name
func (df *DataFrame) getMaxLengthColumn(col string) int {
	strLengths := make([]int, 0, df.length)

	for _, val := range df.data[col] {
		strRepr := fmt.Sprintf("%v", val)
		strLengths = append(strLengths, len(strRepr))
	}
	strLengths = append(strLengths, len(col))

	return helpers.Max(strLengths)
}

// creates the header portion of the dataframe with columns
func (df *DataFrame) createHeader() string {
	sb := strings.Builder{}
	colLengths := make([]int, 0, len(df.columns))

	// calculating header lengths
	for _, col := range df.Columns() {
		colLength := df.getMaxLengthColumn(col)
		colLengths = append(colLengths, colLength)
	}

	// creating the upper and lower bands
	bands := strings.Builder{}
	for i := range df.Columns() {
		bands.WriteString("+" + strings.Repeat("-", colLengths[i]))
	}

	// upper band
	sb.WriteString(bands.String() + "+" + "\n")

	// col name content
	for i, col := range df.Columns() {
		extraSpaces := strings.Repeat(" ", colLengths[i]-len(col))
		sb.WriteString("|" + extraSpaces + col)
	}
	sb.WriteString("|\n")

	// lower band
	sb.WriteString(bands.String() + "+" + "\n")

	return sb.String()
}

func (df *DataFrame) String() string {
	sb := strings.Builder{}
	header := df.createHeader()
	sb.WriteString(header)
	return sb.String()
}
