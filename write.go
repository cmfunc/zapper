package zapper

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Writer struct {
	buf         chan string
	signal      chan struct{}
	f           *os.File
	nextTheTime time.Time
	filepath    string
	tk          time.Duration
	tkStop      chan struct{}
}

func NewWriter(file string, tk time.Duration, cacheMax int64) (w *Writer) {
	// 路径转换（相对路径转绝对路径）
	if !filepath.IsAbs(file) {
		var err error
		file, err = filepath.Abs(file)
		if err != nil {
			panic(err)
		}
	}

	firstTheTime := NextTheTime(time.Now(), tk)
	firstFile := GenLogFilepath(file, firstTheTime)
	f, err := os.OpenFile(firstFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	if cacheMax <= 0 {
		cacheMax = 1024
	}

	w = &Writer{
		buf:         make(chan string, cacheMax),
		signal:      make(chan struct{}),
		f:           f,
		nextTheTime: firstTheTime,
		filepath:    file,
		tk:          tk,
		tkStop:      make(chan struct{}),
	}

	go w.run()
	go w.rotate()

	return
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.buf <- string(p)
	// TODO: whether n,err have been used by zap
	return
}

// 必须在srv.ShutDown()后调用，避免向关闭的通道中发送数据而panic
func (w *Writer) Close() (err error) {
	close(w.buf)
	return
}

func (w *Writer) run() {
	for {
		select {
		case b, ok := <-w.buf:
			if !ok {
				w.tkStop <- struct{}{}
				w.f.Sync()
				w.f.Close()
				return
			}
			fmt.Println("当前日志写入的文件名",w.f.Name()) //TODO: delete
			w.f.WriteString(b)
		case <-w.signal:
			w.f.Sync()
			w.f.Close()
			// replace os.File
			fmt.Println("收到替换信号,当前日志文件:",w.f.Name())
			fmt.Println("收到替换信号,w.nextTheTime:",w.nextTheTime.Format(time.RFC3339))
			newFile := GenLogFilepath(w.filepath, w.nextTheTime)
			fmt.Println("收到替换信号,newFile生成的文件名:",newFile)
			var err error
			w.f, err = os.OpenFile(newFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("收到替换信号,已替换替换的日志文件名:",w.f.Name())
		}
	}
}

func (w *Writer) rotate() {
	for {
		t := time.NewTimer(w.nextTheTime.Sub(time.Now()))
		select {
		case <-w.tkStop:
			return
		case <-t.C:
			// change nextTheTime
			fmt.Println("发送替换日志文件信号时间,NextTheTime前:",w.nextTheTime.Format(time.RFC3339)) // TODO: delete
			w.nextTheTime = NextTheTime(w.nextTheTime, w.tk)
			fmt.Println("发送替换日志文件信号时间,NextTheTime后:",w.nextTheTime.Format(time.RFC3339)) // TODO: delete
			// it's about time to repalce os.File.
			w.signal <- struct{}{}
		}

	}
}
