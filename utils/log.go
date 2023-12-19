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
	Log *logrus.Logger = logrus.New()
)

func init() {
	log_file, err := rotatelogs.New(
		"./log/%Y%m%d_%H:%M:%S.log",
		rotatelogs.WithRotationTime(30*time.Minute),
	)

	if err != nil {
		fmt.Println("[Error] Init log file failed. exiting..")
	}

	Log.SetOutput(io.MultiWriter(log_file, os.Stdout))
	Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
}
