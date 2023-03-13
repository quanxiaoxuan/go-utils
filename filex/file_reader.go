package filex

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件内容
func ReadFile(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if nil != err {
		return "", fmt.Errorf(" %s read file error: %v", filePath, err)
	}
	return string(bytes), nil
}

// 按行读取
func ReadFileLine(path string) ([]string, error) {
	var results []string
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	// 按行处理txt
	reader := bufio.NewReader(file)
	for {
		var line []byte
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		results = append(results, string(line))
	}
	return results, nil
}
