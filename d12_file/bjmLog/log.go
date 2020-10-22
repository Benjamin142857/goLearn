package bjmLog

import (
	"bufio"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota + 1
	INFO
	WARN
	ERROR
)

var LogLevelLabelMap = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type CallInfo struct {
	fileName string
	lineNo   int
	funcName string
}

func getInfo(skip int) (info CallInfo) {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		pFunc := runtime.FuncForPC(pc)
		info.fileName = path.Base(file)
		info.lineNo = line
		info.funcName = pFunc.Name()
	}
	return
}

// LogSource 日志源接口类型 (如 os.File, os.Stdout)
type LogSource interface {
	Write(p []byte) (n int, err error)
}

// Logger 记录器
type Logger struct {
	ls      LogSource
	pWriter *bufio.Writer
	pLevel  *LogLevel
}

// Logger 构建函数
func NewLogger(ls LogSource, level *LogLevel) *Logger {
	w := bufio.NewWriter(ls)
	return &Logger{
		ls:      ls,
		pWriter: w,
		pLevel:  level,
	}
}

// log
func (l *Logger) log(content string, logLevel LogLevel) {
	if *l.pLevel <= logLevel {
		now := time.Now().Format("2006-01-02 15:04:05")
		callInfo := getInfo(3)
		content = strings.Replace(content, "\n", " ", -1)
		content = fmt.Sprintf("[%v|%v|%v:%v %v] %v\n", now, LogLevelLabelMap[logLevel], callInfo.fileName, callInfo.lineNo, callInfo.funcName, content)

		_, _ = l.pWriter.WriteString(content)
		_ = l.pWriter.Flush()
	}

}

// Debug
func (l *Logger) Debug(content string) {
	l.log(content, DEBUG)
}

// Info
func (l *Logger) Info(content string) {
	l.log(content, INFO)
}

// Warn
func (l *Logger) Warn(content string) {
	l.log(content, WARN)
}

// Error
func (l *Logger) Error(content string) {
	l.log(content, ERROR)
}
