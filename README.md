<img src="./logo.png" alt="GoFrames logo">

# Go-Frames

## Introduction

Go-Frames is an ongoing project to build a clone for the python pandas library in Go. This requires an abstract data
structure, that is equivalent to Pandas dataframes and a vast collection of methods that goes along with it. Project is
planned to be extended for machine learning as well by closely following the Sklearn library in Python which would be
compliment to the Go-Frames library. The goal of the project is to get python data scientists to quickly migrate their
code bases to Go for improved performances.

## Basic Usage

### Installation

```bash
go get github.com/pandulaDW/go-frames
```

### Create a Series

A series is the building block of DataFrames. Only a column name and variadic amount of empty interface values are
needed to create a series. Internally the series will be type inferred to be one of Int, Float, Bool, DateTime and
Object(text) types.

```go
package main

import (
	"fmt"

	"github.com/pandulaDW/go-frames/series"
)

func main() {
	s1 := series.NewSeries("col1", 12, 43, 53, 14, 10)
	s2 := series.NewSeries("col2", "foo", "bar", "baz")
	s3 := series.NewSeries("col3", 12.3, 1.43, 4.5)
	s4 := series.NewSeries("col4", true, false, false, true)
	s5 := series.NewSeries("col4", "2010-01-02", "2010-01-02")
}
```

### Reading spreadsheet

The following constitutes the bare to read a spreadsheet document.

```go
package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
```

### Add chart to spreadsheet file

With Excelize chart generation and management is as easy as a few lines of code. You can build charts based on data in
your worksheet or generate charts without any data in your worksheet at all.

<p align="center"><img width="650" src="./test/images/chart.png" alt="Excelize"></p>

```go
package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
```

### Add picture to spreadsheet file

```go
package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet1", "D2", "image.jpg",
		`{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}
	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}
```

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change. XML is
compliant
with [part 1 of the 5th edition of the ECMA-376 Standard for Office Open XML](http://www.ecma-international.org/publications/standards/Ecma-376.htm)
.

## Licenses

This program is under the terms of the BSD 3-Clause License.
See [https://opensource.org/licenses/BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause).

The Excel logo is a trademark of [Microsoft Corporation](https://aka.ms/trademarks-usage). This artwork is an
adaptation.

gopher.{ai,svg,png} was created by [Takuya Ueda](https://twitter.com/tenntenn). Licensed under
the [Creative Commons 3.0 Attributions license](http://creativecommons.org/licenses/by/3.0/).
