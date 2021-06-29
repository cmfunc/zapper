// cron 定时替换io.Writer的落地文件的路径
package zaper

import (
	"io"
	"os"
	"sync"
	"time"
)

type Writer interface {
	io.Writer
	ReciveNewFile(<-chan string)
}

type FileWriter struct {
	File    *os.File
	sync.Mutex
}

func (fw *FileWriter) ReciveNewFile(ch <-chan string) {
	filepath := <-ch
	f, err := os.Create(filepath) //创建文件
	if err != nil {
		panic(err) // 是否应该panic
	}

	fw.Lock()
	defer fw.Unlock()

	fw.File.Sync()
	fw.File.Close()

	fw.File = f
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	fw.Lock()
	defer fw.Unlock()
	n, err = fw.File.Write(p) //写入文件(字节数组)
	return
}

func (fw *FileWriter) CronReplaceWriter(filepath string, writer Writer,tk time.Duration) {

	ticker := time.NewTicker(tk)
	// defer ticker.Stop()
	for range ticker.C {
		// 发送chan 信号通知 io.Writer替换file
		file := filepath + "." + time.Now().Format("2006 15:04:05")
		ch := make(chan string, 1)
		ch <- file
		writer.ReciveNewFile(ch)
	}
}

// NewFileWriter 工厂方法
func NewFileWriter(filepath string) (fw *FileWriter) {
	f, err := os.Create(filepath) //创建文件
	if err != nil {
		panic(err)
	}

	fw = &FileWriter{File: f,}
	go fw.CronReplaceWriter(filepath, fw,time.Hour*1)
	return
}
