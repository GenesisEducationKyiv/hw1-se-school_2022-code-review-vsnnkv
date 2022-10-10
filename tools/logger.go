package tools

import (
	"context"
	"github.com/sirupsen/logrus"
	time "time"
)

var Log = logrus.New()

type LoggerInterface interface {
	LogInfo(msg string)
	LogError(msg string)
	LogDebug(msg string)
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) LogInfo(msg string) {
	Log.Info(msg)
	t := time.Now()
	msg = t.Format("2006-01-02 15:04:05") + " INFO: " + msg
	Publish(context.Background(), msg)
}

func (logger *Logger) LogError(msg string) {
	Log.Error(msg)
	t := time.Now()
	msg = t.Format("2006-01-02 15:04:05") + " ERROR: " + msg
	Publish(context.Background(), msg)

}

func (logger *Logger) LogDebug(msg string) {
	Log.Debug(msg)
	t := time.Now()
	msg = t.Format("2006-01-02 15:04:05") + " DEBUG: " + msg
	Publish(context.Background(), msg)
}
