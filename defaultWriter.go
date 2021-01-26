package golog

import (
	"fmt"
)

type defaultWriter struct {
}

// 实现LogWriter的Write()方法
func (w *defaultWriter) Write(data []byte) (int, error) {
	msg := string(data)
	fmt.Print(msg)
	return len(data), nil
}

func NewDefaultWriter() *defaultWriter {
	return &defaultWriter{}
}
