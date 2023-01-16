package excel

import "github.com/360EntSecGroup-Skylar/excelize"

func ReadFromExcel(filename, sheet string) (rows [][]string, err error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return
	}
	rows = file.GetRows(sheet)
	return
}
