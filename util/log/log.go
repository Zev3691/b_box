package log

import (
	"log"
	"os"
	"time"
)

var loger *log.Logger

func init() {
	file := "./log/" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, "", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
	return
}

func Println(v ...any) {
	loger.Println(v)
}
