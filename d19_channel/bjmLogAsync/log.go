package bjmLogAsync

import (
	"bufio"
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"
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

// LogTask 异步日志任务体
type LogTask struct {
	msg   string
	level LogLevel
	info  CallInfo
}

// LogTask 构造函数
func NewLogTask(msg string, level LogLevel, info CallInfo) *LogTask {
	return &LogTask{msg: msg, level: level, info: info}
}

// Logger 记录器
type Logger struct {
	ls         LogSource
	pWriter    *bufio.Writer
	pLevel     *LogLevel
	ch         chan *LogTask
	chOpLock   sync.Mutex
	chIsClosed bool
}

// Logger 构建函数
func NewLogger(ls LogSource, level *LogLevel) *Logger {
	w := bufio.NewWriter(ls)
	var chOpLock sync.Mutex
	return &Logger{
		ls:       ls,
		pWriter:  w,
		pLevel:   level,
		ch:       make(chan *LogTask, 100),
		chOpLock: chOpLock,
	}
}

// Logger 异步协程开始
func (l *Logger) Run() {
	for pLogTask := range l.ch {
		l.log(pLogTask.msg, pLogTask.level, pLogTask.info)
	}
}

// Logger 异步协程结束
func (l *Logger) Stop() {
	// defer chIsClosed 解锁
	defer l.chOpLock.Unlock()

	// chIsClosed 加锁
	l.chOpLock.Lock()
	if !l.chIsClosed {
		close(l.ch)
		l.chIsClosed = true
	}
}

// log
func (l *Logger) log(content string, logLevel LogLevel, callInfo CallInfo) {
	time.Sleep(time.Millisecond * 1500)
	if *l.pLevel <= logLevel {
		now := time.Now().Format("2006-01-02 15:04:05")
		content = strings.Replace(content, "\n", " ", -1)
		content = fmt.Sprintf("[%v|%v|%v:%v %v] %v\n", now, LogLevelLabelMap[logLevel], callInfo.fileName, callInfo.lineNo, callInfo.funcName, content)

		_, _ = l.pWriter.WriteString(content)
		_ = l.pWriter.Flush()
	}

}

// asyncLog
func (l *Logger) asyncLog(content string, logLevel LogLevel) error {
	// defer chIsClosed 解锁
	defer l.chOpLock.Unlock()

	// chIsClosed 加锁
	l.chOpLock.Lock()
	if !l.chIsClosed {
		l.ch <- NewLogTask(content, logLevel, getInfo(3))
		return nil
	} else {
		return errors.New("the Logger is closed")
	}

}

// Debug
func (l *Logger) Debug(content string) error {
	return l.asyncLog(content, DEBUG)
}

// Info
func (l *Logger) Info(content string) error {
	return l.asyncLog(content, INFO)
}

// Warn
func (l *Logger) Warn(content string) error {
	return l.asyncLog(content, WARN)
}

// Error
func (l *Logger) Error(content string) error {
	return l.asyncLog(content, ERROR)
}
