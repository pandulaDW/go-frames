package series

import "github.com/pandulaDW/go-frames/base"

// GetColumn will return the column type of the series
func (c *Column) GetColumn() *base.Column {
	return &c.column
}

// SetColName will set the column name of the series
func (c *Column) SetColName(colName string) {
	c.column.Name = colName
}

// SetColIndex will set the column index of the series
func (c *Column) SetColIndex(colIndex int) {
	c.column.ColIndex = colIndex
}
