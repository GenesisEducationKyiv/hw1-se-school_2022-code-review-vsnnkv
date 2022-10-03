package tools

import "github.com/sirupsen/logrus"

var Log = logrus.New()

//type Logger interface {
//	LogInfo(msg string)
//	LogError(msg string)
//	LogDebug(msg string)
//}
//
//type LoggerStruct struct {
//}
//
//func NewLogger() *LoggerStruct {
//	return &LoggerStruct{}
//}
//
//func (logger *LoggerStruct) LogInfo(msg string) {
//	log.Info(msg)
//}
//
//func (logger *LoggerStruct) LogError(msg string) {
//	log.Error(msg)
//}
//
//func (logger *LoggerStruct) LogDebug(msg string) {
//	log.Debug(msg)
//}
