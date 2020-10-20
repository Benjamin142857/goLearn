package bjmLog

import (
	"bufio"
	"strings"
)

// LogSource 日志源接口类型 (如 os.File, os.Stdout)
type LogSource interface {
	Write(p []byte) (n int, err error)
}

// Logger 记录器
type Logger struct {
	ls      LogSource
	pWriter *bufio.Writer
}

// Logger 构建函数
func NewLogger(ls LogSource) *Logger {
	w := bufio.NewWriter(ls)
	return &Logger{
		ls:      ls,
		pWriter: w,
	}
}

// info
func (l *Logger) Info(content string) {
	content = strings.Replace(content, "\n", " ", -1)
	content = "time|info:\t" + content + "\n"
	_, _ = l.pWriter.WriteString(content)
	_ = l.pWriter.Flush()
}

// debug
func (l *Logger) Debug(content string) {
	content = strings.Replace(content, "\n", " ", -1)
	content = "time|debug:\t" + content + "\n"
	_, _ = l.pWriter.WriteString(content)
	_ = l.pWriter.Flush()
}

// warn
func (l *Logger) Warn(content string) {
	content = strings.Replace(content, "\n", " ", -1)
	content = "time|warn:\t" + content + "\n"
	_, _ = l.pWriter.WriteString(content)
	_ = l.pWriter.Flush()
}

// error
func (l *Logger) Error(content string) {
	content = strings.Replace(content, "\n", " ", -1)
	content = "time|error:\t" + content + "\n"
	_, _ = l.pWriter.WriteString(content)
	_ = l.pWriter.Flush()
}
