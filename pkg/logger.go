package pkg

import (
    "github.com/sirupsen/logrus"
    "github.com/natefinch/lumberjack"
    "github.com/mattn/go-colorable"
    "os"
	"io"
)

var logLevels = map[string]logrus.Level{
    "error": logrus.ErrorLevel,
    "warn":  logrus.WarnLevel,
    "info":  logrus.InfoLevel,
    "http":  logrus.InfoLevel,
    "debug": logrus.DebugLevel,
}

var Log *logrus.Logger

func InitLogger() {
    Log = logrus.New()

    Log.SetFormatter(&logrus.TextFormatter{
        ForceColors:    true,
        DisableTimestamp: false,
        FullTimestamp:  true,
        TimestampFormat: "2006-01-02 15:04:05",
    })

    Log.SetOutput(colorable.NewColorableStdout())

    Log.SetLevel(logLevels[getLogLevel()])

    Log.SetOutput(io.MultiWriter(
		&lumberjack.Logger{
			Filename:   "./logs/app.log",
			MaxSize:    20,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   true,
		},
		colorable.NewColorableStdout(),
	))
}

func getLogLevel() string {
    env := os.Getenv("NODE_ENV")
    if env == "development" {
        return "debug"
    }
    return "warn"
}