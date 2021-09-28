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
}

func NewWriter(filepath string, tk time.Duration) (w *Writer) {
	firstTheHour := nextTheTime(time.Now(), tk)
	firstFile := genLogFilepath(filepath, firstTheHour)
	f, err := os.Create(firstFile)
	if err != nil {
		panic(err)
	}

	w = &Writer{
		buf:         make(chan []byte, 10240),
		signal:      make(chan struct{}),
		f:           f,
		nextTheTime: firstTheHour,
		filepath:    filepath,
		tk:          tk,
	}

	go w.run()
	go w.notice()

	return 
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.buf <- p
	// TODO: whether n,err have been used by zap
	return
}

func (w Writer) run() {
	for {
		select {
		case b := <-w.buf:
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

func (w Writer) notice() {
	for {

		t := time.NewTimer(w.nextTheTime.Sub(time.Now()))
		<-t.C
		// change nextTheTime
		w.nextTheTime = nextTheTime(w.nextTheTime, w.tk)
		// it's about time to repalce os.File.
		w.signal <- struct{}{}

	}
}
