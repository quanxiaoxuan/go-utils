package excelx

import (
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

// 表头映射
type HeaderList []*Header
type Header struct {
	Key  string
	Name string
}

// 将数据写入excel
func ExcelWriter(excelPath string, headers HeaderList, dataList []map[string]string) error {
	var xlsxFile = xlsx.NewFile()
	var err error
	var sheet *xlsx.Sheet
	sheet, err = xlsxFile.AddSheet("Sheet1")
	if err != nil {
		log.Error("创建excel文件失败：%s", err)
		return err
	}
	// 写入表头
	headerRow := sheet.AddRow()
	for _, header := range headers {
		headerRow.AddCell().Value = header.Name
	}
	// 写入数据
	for _, data := range dataList {
		row := sheet.AddRow()
		for _, header := range headers {
			row.AddCell().Value = data[header.Key]
		}
	}
	//这里从新生成
	err = xlsxFile.Save(excelPath)
	if err != nil {
		return err
	}
	return nil
}
