package log

import (
	"os"
)

// 判断一个文件是否存在
func isFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func saveFile(filePath string, content string) {
	var f *os.File
	var err error
	if isFileExist(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	} else {
		f, err = os.Create(filePath)
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
	f.Sync()
}
