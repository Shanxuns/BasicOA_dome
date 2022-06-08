package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Write(content string) {
	var (
		fileName = "./log.ini"
		file     *os.File
		err      error
	)
	content = time.Now().Format("2006/1/02 15:04:05") + " = " + content + "\n"
	//使用追加模式打开文件
	file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		f, err := os.Create(fileName)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				return
			}
		}(f)
		if err != nil {
			// 创建文件失败处理
			return
		} else {
			_, err = f.Write([]byte(content))
			if err != nil {
				// 写入失败处理
				return
			}
		}
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	//写入文件
	_, _ = file.Write([]byte(content))
	return
}

func Err(app string, text string, err interface{}) {
	var data = "->" + app + ": " + text
	log.Println(data, err)
	Write(data)
}

func App(app string, text string) {
	var data = "->" + app + ": " + text
	log.Println(data)
	Write(data)
}

func Print(text interface{}) {
	fmt.Print("\n", text)
}

func Debug(data ...interface{}) {
	if Mode {
		log.Println(data)
	}
}

func SetMode(mode bool) {
	Mode = mode
}
