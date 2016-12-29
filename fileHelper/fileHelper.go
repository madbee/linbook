package fileHelper

import (
	"io"
	"os"
)

//保存string到文件
func SaveFile(fileName, content string) error {
	var f *os.File
	var err error
	if len(content) > 0 {
		if checkFileExist(fileName) {
			f, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		} else {
			f, err = os.Create(fileName) //创建文件
		}
		_, err = io.WriteString(f, content) //写入文件(字符串)
	}
	return err
}

func checkFileExist(fileName string) (exist bool) {
	exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
