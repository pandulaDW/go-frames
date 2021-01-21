package dataframes

import (
	"fmt"
	"strings"

	"github.com/pandulaDW/go-frames/helpers"
)

// returns the max column size by considering the whole column and the column name
func (df *DataFrame) getMaxLengthColumn(col string) int {
	strLengths := make([]int, 0, df.length)

	for _, val := range df.Data[col].Data {
		strRepr := fmt.Sprintf("%v", val)
		strLengths = append(strLengths, len(strRepr))
	}
	strLengths = append(strLengths, len(col))

	return helpers.MaxIntSlice(strLengths)
}

// creates the header portion of the dataframe with columns
func (df *DataFrame) createHeader(colLengths []int) (string, string) {
	sb := strings.Builder{}

	// creating the upper and lower bands
	band := strings.Builder{}
	for i := range df.Columns() {
		band.WriteString("+" + strings.Repeat("-", colLengths[i]))
	}
	band.WriteByte('+')

	// adding upper band
	sb.WriteString(band.String() + "\n")

	// adding col names content
	for i, col := range df.Columns() {
		extraSpaces := strings.Repeat(" ", colLengths[i]-len(col))
		sb.WriteString("|" + extraSpaces + col)
	}
	sb.WriteString("|\n")

	// adding lower band
	sb.WriteString(band.String() + "\n")

	return sb.String(), band.String()
}

// creates the body portion of the dataframe
func (df *DataFrame) createBody(colLengths []int) string {
	sb := strings.Builder{}

	for i := 0; i < df.length; i++ {
		for colIndex, col := range df.Columns() {
			strRepr := fmt.Sprintf("%v", df.Data[col].Data[i])
			extraSpaces := strings.Repeat(" ", colLengths[colIndex]-len(strRepr))
			sb.WriteString("|" + extraSpaces + strRepr)
		}
		sb.WriteString("|\n")
	}

	return sb.String()
}

func (df *DataFrame) String() string {
	sb := strings.Builder{}
	colLengths := make([]int, 0, len(df.columns))

	// calculating header lengths
	for _, col := range df.Columns() {
		colLength := df.getMaxLengthColumn(col)
		colLengths = append(colLengths, colLength)
	}

	header, band := df.createHeader(colLengths)
	body := df.createBody(colLengths)
	sb.WriteString(header)
	sb.WriteString(body)
	sb.WriteString(band)

	return sb.String()
}
