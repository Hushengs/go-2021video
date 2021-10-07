package ploglib

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func paeseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// Logger
type Logger struct {
	Level LogLevel
}

// NewLog
func NewLog(levelStr string) Logger {
	level, err := paeseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(LogLevel LogLevel) bool {
	return l.Level >= LogLevel
}

func (l Logger) log(lv LogLevel, msg string, arg ...interface{}) {
	if l.enable(lv) {
		format := fmt.Sprintf(msg, arg...)
		fmt.Println(format)
		fileName, funcName, lineNo := getInfo(3)
		now := time.Now().Format("2006-01-02 15:04:06")
		fmt.Printf("%s [%s] [%s %s %d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, format)
	}
}

func (l Logger) Debug(msg string, arg ...interface{}) {
	l.log(DEBUG, msg, arg...)
}

func (l Logger) Trace(msg string, arg ...interface{}) {
	l.log(TRACE, msg, arg...)
}

func (l Logger) Info(msg string, arg ...interface{}) {
	l.log(INFO, msg, arg...)
}

func (l Logger) Warning(msg string, arg ...interface{}) {
	l.log(WARNING, msg, arg...)
}

func (l Logger) Error(msg string, arg ...interface{}) {
	l.log(ERROR, msg, arg...)
}

func (l Logger) Fatal(msg string, arg ...interface{}) {
	l.log(FATAL, msg, arg...)
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEGUB"
	}
}

func getInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime caller failed\n")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return

}
