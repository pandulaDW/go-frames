package dataframes

import (
	"fmt"
	"github.com/pandulaDW/go-frames/base"
	"github.com/pandulaDW/go-frames/helpers"
	"strings"
	"time"
)

// creates the header portion of the dataframe with columns
func (df *DataFrame) createHeader(colLengths []int, isCustomIndex bool) (string, string) {
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

	// adding index col header if index is custom
	if isCustomIndex {
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
	}

	// adding lower band
	sb.WriteString(band.String() + "\n")

	return sb.String(), band.String()
}

// createRowString creates a string representation of a row
func (df *DataFrame) createRowString(i int, colLengths []int, sb *strings.Builder) {
	for _, col := range df.columns {
		var strRepr string
		val := df.data[col.Name].Data[i]
		if col.Dtype == base.DateTime && helpers.IsTimeSet(val.(time.Time)) {
			strRepr = val.(time.Time).Format("2006-01-02")
		} else {
			strRepr = fmt.Sprintf("%v", val)
		}
		extraSpaces := strings.Repeat(" ", colLengths[col.ColIndex]-len(strRepr))
		sb.WriteString("|" + extraSpaces + strRepr)
	}
}

// creates the body portion of the dataframe
func (df *DataFrame) createBody(colLengths []int) string {
	sb := strings.Builder{}

	for i := 0; i < df.length; i++ {
		df.createRowString(i, colLengths, &sb)
		sb.WriteString("|\n")
	}

	return sb.String()
}

// creates the body portion for large dataframes
func (df *DataFrame) createBodyLarge(colLengths []int) string {
	sb := strings.Builder{}

	for i := 0; i < 5; i++ {
		df.createRowString(i, colLengths, &sb)
		sb.WriteString("|\n")
	}

	for _, col := range df.columns {
		strRepr := "..."
		extraSpaces := strings.Repeat(" ", colLengths[col.ColIndex]-len(strRepr))
		sb.WriteString("|" + extraSpaces + strRepr)
	}
	sb.WriteString("|\n")

	for i := df.length - 5; i < df.length; i++ {
		df.createRowString(i, colLengths, &sb)
		sb.WriteString("|\n")
	}

	return sb.String()
}

func (df *DataFrame) String() string {
	sb := strings.Builder{}

	// adding index column as the first column
	indexColName := df.Index.Data.GetColumn().Name
	cols := append([]string{indexColName}, df.Columns()...)
	copiedDF := df.ShallowCopy().WithColumn(df.Index.Data).ResetColumns(cols)

	colLengths := make([]int, 0, len(copiedDF.columns))

	// calculating header lengths
	for _, col := range copiedDF.Columns() {
		colLength := copiedDF.data[col].GetMaxLength()
		colLengths = append(colLengths, colLength)
	}

	header, band := copiedDF.createHeader(colLengths, df.Index.IsCustom)
	var body string
	if df.length < 100 {
		body = copiedDF.createBody(colLengths)
	} else {
		body = copiedDF.createBodyLarge(colLengths)
	}
	sb.WriteString(header)
	sb.WriteString(body)
	sb.WriteString(band)

	return sb.String()
}
