package ploglib

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := paeseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic("open log failed")
	}
	return f1
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log error file failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return f.Level >= LogLevel
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInof, err := f.fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	return fileInof.Size() >= f.maxFileSize
}

func (f *FileLogger) log(lv LogLevel, msg string, arg ...interface{}) {
	if f.enable(lv) {
		format := fmt.Sprintf(msg, arg...)
		fmt.Println(format)

		fileName, funcName, lineNo := getInfo(3)
		now := time.Now().Format("2006-01-02 15:04:06")
		if f.checkSize(f.fileObj) {
			fileObj, err := f.splitFile(f.fileObj)
			if err != nil {
				fmt.Printf("splitFile file failed,err:%v\n", err)
			}
			f.fileObj = fileObj
		}
		fmt.Fprintf(f.fileObj, "%s [%s] [%s %s %d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, format)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				fileObj, err := f.splitFile(f.errFileObj)
				if err != nil {
					fmt.Printf("splitFile file failed,err:%v\n", err)
				}
				f.fileObj = fileObj
			}
			fmt.Fprintf(f.errFileObj, "%s [%s] [%s %s %d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, format)
		}
	}
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割
	//1.关闭当前的日志文件
	//2.备份一下，rename
	//3.打开一个新的日志文件
	//4.将打开的新日志文件对象赋值给f.fileObj
	f.fileObj.Close()
	nowStr := time.Now().Format("20060102150405000")
	logName := path.Join(f.filePath, f.fileName)
	nowlog := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(logName, nowlog)
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open new log file failed,err:%v\n", err)
		return file, err
	}
	return fileObj, nil
}

func (f *FileLogger) Debug(msg string, arg ...interface{}) {
	f.log(DEBUG, msg, arg...)
}

func (f *FileLogger) Trace(msg string, arg ...interface{}) {
	f.log(TRACE, msg, arg...)
}

func (f *FileLogger) Info(msg string, arg ...interface{}) {
	f.log(INFO, msg, arg...)
}

func (f *FileLogger) Warning(msg string, arg ...interface{}) {
	f.log(WARNING, msg, arg...)
}

func (f *FileLogger) Error(msg string, arg ...interface{}) {
	f.log(ERROR, msg, arg...)
}

func (f *FileLogger) Fatal(msg string, arg ...interface{}) {
	f.log(FATAL, msg, arg...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
