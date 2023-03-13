package filex

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
)

// 写入文件
func WriteFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(content)
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// 数组按行写入文件
func WriteFileLine(path string, arrays []string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	writer := bufio.NewWriter(file)
	for _, line := range arrays {
		_, _ = writer.WriteString(line)
		_, _ = writer.WriteString("\n")
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// 写入json文件
func WriteJson(path string, obj interface{}) error {
	jsonByte, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, jsonByte, 0777)
	if err != nil {
		return err
	}
	return nil
}

// 写入csv文件
func WriteCSV(path string, data [][]string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ','
	csvWriter.UseCRLF = true
	err = csvWriter.WriteAll(data)
	if err != nil {
		return err
	}
	csvWriter.Flush()
	return nil
}
