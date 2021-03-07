// include tests for dataframes where circular dependency problem occurs
package main

import (
	"github.com/pandulaDW/go-frames/ioread"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDisplay(t *testing.T) {
	// assert that display string for large dataframes works as expected
	expected := `
+----+-----------+------+-------+-------+-----+--------+----+
|    |       date|  open|   high|    low|close|  volume|Name|
+----+-----------+------+-------+-------+-----+--------+----+
|   0| 2013-02-08| 45.07|  45.35|     45|45.08| 1824755|   A|
|   1| 2013-02-11| 45.17|  45.18|  44.45| 44.6| 2915405|   A|
|   2| 2013-02-12| 44.81|  44.95|   44.5|44.62| 2373731|   A|
|   3| 2013-02-13| 44.81|  45.24|  44.68|44.75| 2052338|   A|
|   4| 2013-02-14| 44.72|  44.78|  44.36|44.58| 3826245|   A|
| ...|        ...|   ...|    ...|    ...|  ...|     ...| ...|
|1254| 2018-02-01| 73.18|  73.78|  72.51|72.83| 2008177|   A|
|1255| 2018-02-02| 72.32|  72.76|  71.22|71.25| 1955697|   A|
|1256| 2018-02-05| 70.86|  71.48|  68.18|68.22| 2860726|   A|
|1257| 2018-02-06| 66.96|  68.83|  66.13|68.45| 4121239|   A|
|1258| 2018-02-07| 68.19| 69.085| 67.905|68.06| 1988391|   A|
+----+-----------+------+-------+-------+-----+--------+----+
`
	expected = strings.TrimSpace(expected)
	df, _ := ioread.ReadCSV(ioread.CsvOptions{Path: "data/A_data.csv"})
	_ = df.Data["date"].CastAsTime("2006-01-02")
	assert.Equal(t, expected, df.String())
}
