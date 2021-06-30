// cron 定时替换io.Writer的落地文件的路径
// 使用sync/atomic替换mutex
package zaper

import (
	"io"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

type Writer interface {
	io.Writer
	ReciveNewFile(<-chan string)
}

type FileWriter struct {
	atomic.Value
}

func (fw *FileWriter) ReciveNewFile(ch <-chan string) {
	filepath := <-ch
	f, _ := os.Create(filepath)

	x := fw.Load()
	if v, ok := x.(*os.File); ok {
		v.Sync()
		v.Close()
	}

	fw.Store(f)
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	x := fw.Load()
	if v, ok := x.(*os.File); ok {
		n, err = v.Write(p)
	}

	return
}

func (fw *FileWriter) CronReplaceWriter(filepath string, tk time.Duration) {

	ticker := time.NewTicker(tk)
	// defer ticker.Stop()
	for range ticker.C {
		// 发送chan 信号通知 io.Writer替换file
		file := strings.Replace(filepath, ".log", time.Now().Format("2006 15:04:05")+".log", 1)
		ch := make(chan string, 1)
		ch <- file
		fw.ReciveNewFile(ch)
	}
}

// NewFileWriter 工厂方法
func NewFileWriter(filepath string) (fw *FileWriter) {
	f, _ := os.Create(filepath) //创建文件

	fw = &FileWriter{}
	fw.Store(f)
	go fw.CronReplaceWriter(filepath, time.Second*5)
	return
}
