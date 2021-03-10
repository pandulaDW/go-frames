package series

import (
	"github.com/pandulaDW/go-frames/base"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeries_GetColumn(t *testing.T) {
	newCol := base.Column{Name: "col", Dtype: base.Float64}
	col := Column{column: newCol}
	// assert that Column returned correctly
	assert.Equal(t, col.GetColumn(), newCol)
}

func TestSeries_SetColName(t *testing.T) {
	newCol := base.Column{Name: "col", Dtype: base.Float64}
	col := Column{column: newCol}
	col.SetColName("newColName")
	// assert that Column name is set correctly
	assert.Equal(t, "newColName", col.column.Name)
}

func TestSeries_SetColIndex(t *testing.T) {
	newCol := base.Column{Name: "col", Dtype: base.Float64}
	col := Column{column: newCol}
	col.SetColIndex(3)
	// assert that Column index set correctly
	assert.Equal(t, 3, col.column.ColIndex)
}
