package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logEnt *logrus.Logger

func Init() {
	logEnt = logrus.New()
	name := time.Now().Format("20060102")
	fd, err := os.OpenFile("./log/"+name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic("创建日志文件失败")
	}
	logEnt.SetOutput(fd)
	logEnt.Printf("\n\n\n\n%s", "----------------------------------------")
}

func Println(format string, args ...interface{}) {
	logEnt.Infof(format, args)
}
