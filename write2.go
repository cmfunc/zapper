package zapper

import (
	"os"
	"time"
)

type Writer struct {
	buf         chan []byte
	signal      chan struct{}
	f           *os.File
	nextTheTime time.Time
	filepath    string
	tk          time.Duration
	tkStop      chan struct{}
}

func NewWriter(filepath string, tk time.Duration, cacheMax int64) (w *Writer) {
	firstTheHour := nextTheTime(time.Now(), tk)
	firstFile := genLogFilepath(filepath, firstTheHour)
	f, err := os.Create(firstFile)
	if err != nil {
		panic(err)
	}

	if cacheMax == 0 {
		cacheMax = 1024
	}

	w = &Writer{
		buf:         make(chan []byte, cacheMax),
		signal:      make(chan struct{}),
		f:           f,
		nextTheTime: firstTheHour,
		filepath:    filepath,
		tk:          tk,
		tkStop:      make(chan struct{}),
	}

	go w.run()
	go w.rotate()

	return
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.buf <- p
	// TODO: whether n,err have been used by zap
	return
}

// 必须在srv.ShutDown()后调用，避免向关闭的通道中发送数据而panic
func (w Writer) Close() (err error) {
	close(w.buf)
	return
}

func (w Writer) run() {
	for {
		select {
		case b, ok := <-w.buf:
			if !ok {
				w.tkStop <- struct{}{}
				w.f.Sync()
				w.f.Close()
				return
			}
			w.f.Write(b)
		case <-w.signal:
			w.f.Sync()
			w.f.Close()
			// replace os.File
			firstFile := genLogFilepath(w.filepath, w.nextTheTime)
			w.f, _ = os.Create(firstFile)
		}
	}
}

func (w Writer) rotate() {
	for {
		t := time.NewTimer(w.nextTheTime.Sub(time.Now()))
		select {
		case <-w.tkStop:
			return
		case <-t.C:
			// change nextTheTime
			w.nextTheTime = nextTheTime(w.nextTheTime, w.tk)
			// it's about time to repalce os.File.
			w.signal <- struct{}{}
		}

	}
}
