package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	categories := []string{"ID", "Apple", "Orange"} //表头
	values := map[string]interface{}{"A2": 2, "B2": "3", "C2": 3, "A3": 5, "B3": 2, "C3": 4, "A4": 6, "B4": 7, "C4": 8}
	xlsx := excelize.NewFile()
	d := ""
	for k, v := range categories {
		// cl := fmt.Sprintf("%d", (k + 1))
		cl := "1"
		fmt.Println(cl)
		if k == 0 {
			d = "A" + cl
		} else if k == 1 {
			d = "B" + cl
		} else if k == 2 {
			d = "C" + cl
		} else if k == 3 {
			d = "D" + cl
		}
		fmt.Println(d, v)
		xlsx.SetCellValue("Sheet1", d, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	// xlsx.AddChart("Sheet1", "E1", `{"type":"bar3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Line Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
