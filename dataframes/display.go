package dataframes

import (
	"fmt"
	"strings"
)

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
		if i == 0 {
			spaces := strings.Repeat(" ", colLengths[i])
			sb.WriteString("|" + spaces)
			continue
		}
		extraSpaces := strings.Repeat(" ", colLengths[i]-len(col))
		sb.WriteString("|" + extraSpaces + col)
	}
	sb.WriteString("|\n")

	// adding index col header
	for i, col := range df.Columns() {
		if i == 0 {
			extraSpaces := strings.Repeat(" ", colLengths[i]-len(col))
			sb.WriteString("|" + extraSpaces + col)
			continue
		}
		spaces := strings.Repeat(" ", colLengths[i])
		sb.WriteString("|" + spaces)
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

	// adding index column as the first column
	indexColName := df.Index.Data.GetColumn().Name
	cols := append([]string{indexColName}, df.Columns()...)
	copiedDF := df.ShallowCopy().AddColumn(df.Index.Data).ResetColumns(cols)

	colLengths := make([]int, 0, len(copiedDF.columns))

	// calculating header lengths
	for _, col := range copiedDF.Columns() {
		colLength := copiedDF.Data[col].GetMaxLength()
		colLengths = append(colLengths, colLength)
	}

	header, band := copiedDF.createHeader(colLengths)
	body := copiedDF.createBody(colLengths)
	sb.WriteString(header)
	sb.WriteString(body)
	sb.WriteString(band)

	return sb.String()
}
