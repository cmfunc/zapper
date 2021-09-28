// cron 定时替换io.Writer的落地文件的路径
// 使用sync/atomic替换mutex
package zapper

import (
	"os"
	"strings"
	"time"
)

type FileWriter struct {
	ch chan *os.File
}

// 实现io.Writer接口
func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f := <-fw.ch
	n, err = f.Write(p)
	fw.ch <- f
	return
}

// NewFileWriter
// filepath is default log filepath,
// tk is time when next log file will be generated cyclical.
func NewFileWriter(filepath string, tk time.Duration) (fw *FileWriter) {
	firstTheHour := nextTheTime(time.Now(), tk)
	firstFile := genLogFilepath(filepath, firstTheHour)
	f, _ := os.Create(firstFile)

	fw = &FileWriter{ch: make(chan *os.File, 1)}
	fw.ch <- f

	go func() {
		for {
			now := firstTheHour.Add(tk)
			next := nextTheTime(now, tk)
			t := time.NewTimer(next.Sub(now))
			<-t.C
			// it's about time, to rotate log file.
			file := genLogFilepath(filepath, next)
			f, _ := os.Create(file)
			// ReciveNewFile accept new *os.File
			// with a lock by channel
			oldFile := <-fw.ch

			go func(f *os.File) {
				// TODO: 控制goroutine的生命周期,logger.Close时,必须等待goroutine关闭;
				f.Sync()
				f.Close()
			}(oldFile)

			fw.ch <- f
		}
	}()
	return
}

func genLogFilepath(filepath string, nextTheTime time.Time) string {
	fileSlice := strings.Split(filepath, ".")
	return strings.Join([]string{fileSlice[0], nextTheTime.Format("2006 15:04:05"), ".log"}, "")
}

func nextTheTime(t time.Time, tk time.Duration) time.Time {
	t = t.Add(tk)
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}
