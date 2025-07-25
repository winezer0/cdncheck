package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadJsonToMap 读取并解析一个 JSON 文件，返回 map[string]interface{}
func ReadJsonToMap(filePath string) (map[string]interface{}, error) {
	// 打开文件
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取文件: %w", err)
	}

	// 解析 JSON 数据
	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析失败: %w", err)
	}

	return data, nil
}

func ReadJsonToStruct(filePath string, v interface{}) error {
	file, err := os.ReadFile(filePath)
	if err != nil {

		return fmt.Errorf("无法读取文件: %w", err)
	}
	err = json.Unmarshal(file, v)
	return err
}
