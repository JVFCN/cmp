package Methods

import (
	"fmt"
	"os"
)

var TEMPDIR = os.TempDir() + "\\cmp\\"

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()
}

func EnsureDir(path string) error {
	// 检查目录是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 如果目录不存在，则创建目录
		err := os.MkdirAll(path, os.ModePerm) // os.ModePerm 为 0777 权限
		if err != nil {
			return fmt.Errorf("创建目录失败: %v", err)
		}
	} else if err != nil {
		// 处理其他错误（如权限问题）
		return fmt.Errorf("检查目录时发生错误: %v", err)
	}
	return nil
}
