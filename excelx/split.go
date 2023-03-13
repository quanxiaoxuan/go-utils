package excelx

import (
	"fmt"
	"strings"

	"github.com/quanxiaoxuan/go-utils/stringx"
	"github.com/tealeg/xlsx"
)

type SheetInfoList []*SheetInfo
type SheetInfo struct {
	SheetName string `json:"sheetName"`
	StartRow  int    `json:"startRow"`
	EndRow    int    `json:"endRow"`
}

// 在目标sheet页中，根据【具有合并单元格的行】进行横向拆分
// 取合【并单元格的值】作为拆分出的新sheet名
// excelPath:目标excel
// sheetName:目标sheet页,默认取第一个sheet页
func ExcelSplit(excelPath, sheetName string) (string, error) {
	var err error
	// 读取excel
	var xlsxFile *xlsx.File
	xlsxFile, err = xlsx.OpenFile(excelPath)
	if err != nil {
		return excelPath, err
	}
	// 读取目标sheet
	var theSheet *xlsx.Sheet
	if xlsxFile.Sheet[sheetName] == nil {
		sheetName = xlsxFile.Sheets[0].Name
	}
	for _, sheet := range xlsxFile.Sheets {
		if sheet.Name == sheetName {
			theSheet = sheet
		}
	}
	// 新增sheet页
	var NewExcelFile = xlsx.NewFile()
	var addSheetList SheetInfoList
	for rowNo, rowData := range theSheet.Rows {
		// 如果是合并单元格
		if rowData.Cells == nil || len(rowData.Cells) == 0 {
			continue
		} else if rowData.Cells[0].HMerge > 0 {
			addSheetName := rowData.Cells[0].Value
			addSheetName, _ = stringx.SplitByLast(addSheetName, `"`, false)
			_, addSheetName = stringx.SplitByLast(addSheetName, `"`, false)
			addSheetName = strings.ReplaceAll(addSheetName, "ARCHIVE_", "")
			addSheetName = strings.ReplaceAll(addSheetName, "WORKFLOW_PERSON_LIST_03_PERSON_LINK_LIB", "WORKFLOW_PERSON_LINK_LIB")
			_, err = NewExcelFile.AddSheet(addSheetName)
			fmt.Println(addSheetName, "===", err)
			addSheetList = append(addSheetList, &SheetInfo{addSheetName, rowNo, 0})
		} else {
			continue
		}
	}
	for i, item := range addSheetList {
		item.StartRow = item.StartRow + 1
		if i < len(addSheetList)-1 {
			item.EndRow = addSheetList[i+1].StartRow - 2
		} else {
			item.EndRow = len(theSheet.Rows) - 1
		}
	}
	for _, item := range addSheetList {
		if NewExcelFile.Sheet[item.SheetName] != nil {
			for rowNo, rowData := range theSheet.Rows {
				if rowNo >= item.StartRow && rowNo <= item.EndRow {
					row := NewExcelFile.Sheet[item.SheetName].AddRow()
					for _, cell := range rowData.Cells {
						row.AddCell().Value = cell.Value
					}
				}
			}
		}
	}
	dir, excelName := stringx.SplitByLast(excelPath, "\\", true)
	excelName, _ = stringx.SplitByFirst(excelName, ".", false)
	newExcelPath := dir + excelName + "_split.xlsx"
	err = NewExcelFile.Save(newExcelPath)
	if err != nil {
		return excelPath, err
	}
	return newExcelPath, nil
}
