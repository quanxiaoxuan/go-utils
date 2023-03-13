package excelx

import (
	"github.com/tealeg/xlsx"
)

// excelPath:目标excel
// sheetName:目标sheet页,默认取第一个sheet页
func ExcelReader(excelPath, sheetName string, headerMap map[string]string) ([]map[string]string, error) {
	var resultMapList []map[string]string
	var err error
	// 读取excel
	var xlsxFile *xlsx.File
	xlsxFile, err = xlsx.OpenFile(excelPath)
	if err != nil {
		return nil, err
	}
	// 读取目标sheet
	var theSheet *xlsx.Sheet
	if xlsxFile.Sheet[sheetName] == nil {
		theSheet = xlsxFile.Sheets[0]
	} else {
		theSheet = xlsxFile.Sheet[sheetName]
	}
	// 读取表头
	var headers []string
	for _, cell := range theSheet.Rows[0].Cells {
		header := headerMap[cell.Value]
		if header == "" {
			header = cell.Value
		}
		headers = append(headers, header)
	}
	// 遍历excel(x:横向坐标，y:纵向坐标)
	for y, row := range theSheet.Rows {
		if y > 0 {
			var rowMap = make(map[string]string)
			for x, cell := range row.Cells {
				if x >= len(headers) {
					break
				}
				rowMap[headers[x]] = cell.Value
			}
			resultMapList = append(resultMapList, rowMap)
		}
	}
	return resultMapList, nil
}
