package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var (
	INFO  *log.Logger
	ERROR *log.Logger
	DEBUG *log.Logger
)

func Log() {
	var (
		file    *os.File
		errFile *os.File
	)
	for {
		prefix := time.Now().Format("2006-01-02")
		if file != nil && path.Base(file.Name()) != prefix+".log" {
			fmt.Println("关闭文件")
			file.Close()
		}
		fileName := "logs/" + prefix + ".log"
		file, _ = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		errFile, _ = os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		// 在日志中包含日期，时间，文件名和行号
		INFO = log.New(file, "INFO：", log.Ldate|log.Ltime|log.Lshortfile)
		DEBUG = log.New(file, "DEBUG：", log.Ldate|log.Ltime|log.Lshortfile)
		ERROR = log.New(errFile, "ERROR：", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
