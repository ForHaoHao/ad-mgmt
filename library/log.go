package library

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	Log.SetOutput(&lumberjack.Logger{
		Filename: "app.log",
		MaxSize:  10,
		MaxAge:   28,
		Compress: true,
	})

	Log.SetLevel(logrus.InfoLevel)
}
