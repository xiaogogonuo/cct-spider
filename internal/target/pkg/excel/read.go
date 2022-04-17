package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadExcel(path, sheet string) [][]string {
	file, err := excelize.OpenFile(path)
	if err != nil {
		panic(err)
	}
	return file.GetRows(sheet)
}
