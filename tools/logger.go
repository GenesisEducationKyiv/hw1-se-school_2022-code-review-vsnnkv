package tools

import (
	"github.com/sirupsen/logrus"
	time "time"
)

var Log = logrus.New()

type LoggerStruct struct {
}

func NewLogger() *LoggerStruct {
	return &LoggerStruct{}
}

func (logger *LoggerStruct) LogInfo(msg string) string {
	Log.Info(msg)
	t := time.Now()

	return t.Format("2006-01-02 15:04:05") + " INFO: " + msg
}

func (logger *LoggerStruct) LogError(msg string) string {
	Log.Error(msg)
	t := time.Now()

	return t.Format("2006-01-02 15:04:05") + " ERROR: " + msg

}

func (logger *LoggerStruct) LogDebug(msg string) string {
	Log.Debug(msg)
	t := time.Now()

	return t.Format("2006-01-02 15:04:05") + " DEBUG: " + msg
}
