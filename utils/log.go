package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger = logrus.New()
	Log Logger         = Logger{
		Info: func(msg string, fields map[string]interface{}) {
			log.WithFields(fields).Info(msg)
		},
		Error: func(msg string, fields error) {
			log.WithFields(logrus.Fields{"err": fields}).Error(msg)
		},
		ErrorP: func(msg string, fields map[string]interface{}) {
			log.WithFields(fields).Error(msg)
		},
	}
)

func init() {
	log_file, err := rotatelogs.New(
		"./log/%Y%m%d_%H:%M:%S.log",
		rotatelogs.WithRotationTime(30*time.Minute),
	)

	if err != nil {
		fmt.Println("[Error] Init log file failed. exiting..")
	}

	log.SetOutput(io.MultiWriter(log_file, os.Stdout))
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
}

type Logger struct {
	Info   func(string, map[string]interface{})
	Error  func(string, error)
	ErrorP func(string, map[string]interface{})
}
