package jsonx

import (
	"encoding/json"
	"os"
)

// 将json字符串转为interface{}
func JsonToInterface(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return err
	}
	return nil
}

// 结构体转Json
func InterfaceToJson(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 结构写入json文件
func structWriteJsonFile(path string, obj interface{}) {
	jsonByte, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return
	}
	err = os.WriteFile(path, jsonByte, 0777)
	if err != nil {
		return
	}
}
