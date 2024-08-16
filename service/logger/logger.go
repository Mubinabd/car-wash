package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logFile *os.File

func InitLog() {
	// var err error
	// logFile, err = os.OpenFile("./logs/carwash.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
}

func SetOutput(output io.Writer) {
	logrus.SetOutput(output)
}

func Debug(msg ...interface{}) {
	logrus.Debugln(msg...)
}

func Info(msg ...interface{}) {
	logrus.Infoln(msg...)
}

func Warn(msg ...interface{}) {
	logrus.Warnln(msg...)
}

func Error(msg ...interface{}) {
	logrus.Errorln(msg...)
}

func Fatal(msg ...interface{}) {
	logrus.Fatalln(msg...)
}
