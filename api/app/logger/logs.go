package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	// gas "github.com/firstrow/goautosocket"

	"github.com/sirupsen/logrus"
)

//const (
//    LogsDirpath = "logs"
//)

type LogDir struct {
	LogDirectory string
}

//func New() *LogDir {
//    err := os.Mkdir(LogsDirpath, 0666)
//    if err != nil {
//        return nil
//    }
//    return &LogDir{
//        LogDirectory: LogsDirpath,
//    }
//}

func SetLogFile() *os.File {
	projectPath := os.Getenv("PROJECT_SOURCE_CODE")
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	fileNameAndPath := projectPath + "/app/logs/" + fileName
	fmt.Printf("Log file " + fileNameAndPath)

	filePath, _ := os.OpenFile(fileNameAndPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return filePath
}

func (l *LogDir) Info() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Warning() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Error() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Fatal() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLog() *logrus.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})

	// logstash := os.Getenv("LOGSTASH_ENDPOINT")

	logr := logrus.New()

	// conn, err := gas.Dial("tcp", logstash)
	// if err != nil {
	// 	log.Printf("error here %s",err.Error())
	// 	return logrus.StandardLogger()
	// }

	// hook, _ := logrustash.NewHookWithConn(conn, "Identity Service")
	// logr.Hooks.Add(hook)
	return logr.WithFields(logrus.Fields{"method": "main"}).Logger
}
