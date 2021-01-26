package golog

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
)

var loger *Logger

// 声明日志写入器接口

// 日志器
type Logger struct {
	// 这个日志器用到的日志写入器
	writerList []io.Writer
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer io.Writer) {
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的数据写入日志
func (l *Logger) Log(data []byte) {
	// 遍历所有注册的写入器
	for _, writer := range l.writerList {
		// 将日志输出到每一个写入器中
		writer.Write(data)
	}
}

func (l *Logger) LogF(format string, args ...interface{}) {

	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)

	//fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	data := fmt.Sprintf("%s:%d:%s:  %s\n", filename, line, funcname, fmt.Sprintf(format, args...))

	l.Log([]byte(data))
}

// 创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}

func init() {
	if loger == nil {
		loger = NewLogger()
	}

	// 注册默认的标准输出
	writer := NewDefaultWriter()
	loger.RegisterWriter(writer)
}

func RegisterWriter(writer io.Writer) {
	if loger == nil {
		return
	}

	loger.writerList = append(loger.writerList, writer)
}

func Debug(format string, args ...interface{}) {
	if loger == nil {
		return
	}

	loger.LogF(format, args...)
}
