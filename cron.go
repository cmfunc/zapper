// cron 定时替换io.Writer的落地文件的路径
// 使用sync/atomic替换mutex
package zaper

import (
	"io"
	"os"
	"strings"
	"time"
)

type Writer interface {
	io.Writer
	ReciveNewFile(<-chan string)
}

type FileWriter struct {
	ch chan *os.File
}

// ReciveNewFile accept new *os.File
// with a lock by channel
func (fw *FileWriter) ReciveNewFile(ch <-chan string) {
	filepath := <-ch
	f, _ := os.Create(filepath)

	oldWiter := <-fw.ch
	oldWiter.Sync()
	oldWiter.Close()
	fw.ch <- f
}

// Write io.Writer imp
func (fw *FileWriter) Write(p []byte) (n int, err error) {

	f := <-fw.ch
	n, err = f.Write(p)
	fw.ch <- f

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
func NewFileWriter(filepath string, syncCycle time.Duration) (fw *FileWriter) {
	firstFile := strings.Replace(filepath, ".log", time.Now().Format("2006 15:04:05")+".log", 1)
	f, _ := os.Create(firstFile)

	fw = &FileWriter{ch: make(chan *os.File, 1)}
	fw.ch <- f

	go fw.CronReplaceWriter(filepath, syncCycle)
	return
}
