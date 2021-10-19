package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	rl, err := rotatelogs.New("test.log_%Y%m%d%H%M")
	if err != nil {
		panic(err)
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	mw := io.MultiWriter(os.Stdout, rl)
	logrus.SetOutput(mw)

	logrus.Info("hello world")

}
