package zapper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var defaultWriter *Writer

// SetDefaultWriter
// 设置包内默认的日志Writer，便于日志Sync()时，对Writer执行收尾工作
func SetDefaultWriter(w *Writer) { defaultWriter = w }

// Writer 日志输出的实现
// buf 为写日志的缓冲
// signal 为切换日志文件的信号
// f为当前输出文件指针
// nextTime 保存下一个整点时间
// filepath 为默认的日志文件命名
// tk 为整点间隔单位
// tkSStop 为日志切割停止信号
type Writer struct {
	buf          chan string   // 日志数据缓冲
	rotateSignal chan struct{} // 日志切割信号
	f            *os.File      // 当前落盘日志文件指针
	nextTheTime  time.Time     // 当前日志所在整点时间
	filepath     string        // 默认日志文件地址
	tk           time.Duration // 日志切割间隔整点时间
	tkStop       chan struct{} // 日志切割停止信号
	syncFileBuf  chan *os.File // 需要异步关闭的文件指针
}

// Clone 复制一个新的Writer
func (w *Writer) Clone() *Writer {
	copy := *w
	return &copy
}

// WithOptions 对Writer进行可选配置
func (w *Writer) WithOptions(opts ...WriterOption) *Writer {
	c := w.Clone()
	for _, opt := range opts {
		opt.Apply(c)
	}
	return c
}

// WriterOption Writer的可选配置接口
type WriterOption interface {
	Apply(*Writer)
}

// WriterOption 可选配置方法实例
type WriterOptionFunc func(*Writer)

// Apply WriterOption可选配置方法实例方法
func (f WriterOptionFunc) Apply(w *Writer) {
	f(w)
}

// NewWriter Writer的工厂方法
// file 文件名
// tk 日志切割的间隔整点时间单位
// cacheMax 缓存的最大值
// Writer 的配置项做成Options方式处理
func NewWriter(file string, tk time.Duration, cacheMax int64) (w *Writer) {
	// 路径转换（相对路径转绝对路径）
	var err error
	file, err = filepath.Abs(file) //绝对路径会直接返回绝对路径
	if err != nil {
		panic(err)
	}

	firstTheTime := NextTheTime(time.Now(), tk)
	firstFile := GenLogFilepath(file, firstTheTime, tk)
	f, err := os.OpenFile(firstFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	if cacheMax <= 0 {
		cacheMax = 1024
	}

	w = &Writer{
		buf:          make(chan string, cacheMax),
		rotateSignal: make(chan struct{}),
		f:            f,
		nextTheTime:  firstTheTime,
		filepath:     file,
		tk:           tk,
		tkStop:       make(chan struct{}),
		syncFileBuf:  make(chan *os.File, 128),
	}

	go w.run()
	go w.rotate()
	go w.syncSync()

	return
}

// Write
// io.Writer interface的实现
func (w *Writer) Write(p []byte) (n int, err error) {
	w.buf <- string(p)
	return
}

// Close Writer的收尾方法
// 必须在srv.ShutDown()后调用，避免向关闭的通道中发送数据而panic
func (w *Writer) Close() (err error) {
	close(w.buf)
	return
}

// run Writer的实际处理方法
// 会对日志切割型号、日志罗盘数据进行处理
func (w *Writer) run() {
	for {
		select {
		case b, ok := <-w.buf:
			if !ok {
				w.tkStop <- struct{}{}
				w.f.Sync() // TODO: 检查logger.Sync()
				w.f.Close()
				return
			}
			// check json repeated key
			fmt.Println("准备写入文件的日志内容:", b)
			bm := make(map[string]interface{}, 0)
			err := json.Unmarshal([]byte(b), &bm)
			if err != nil {
				fmt.Println(err)
			}

			bb, err := json.Marshal(bm)
			if err != nil {
				fmt.Println(err)
			}
			// w.f.WriteString(b)
			w.f.WriteString(string(bb) + "\n")
		case <-w.rotateSignal:
			// 异步化处理文件同步、关闭
			w.syncFileBuf <- w.f
			// replace os.File
			newFile := GenLogFilepath(w.filepath, w.nextTheTime, w.tk)
			var err error
			w.f, err = os.OpenFile(newFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// rotate
// 定时发送日志切割信号
// 监听切割终止信号
func (w *Writer) rotate() {
	for {
		t := time.NewTimer(w.nextTheTime.Sub(time.Now()))
		select {
		case <-w.tkStop:
			close(w.syncFileBuf) // 通知异步任务关闭
			return
		case <-t.C:
			// change nextTheTime
			w.nextTheTime = NextTheTime(w.nextTheTime, w.tk)
			// it's about time to repalce os.File.
			w.rotateSignal <- struct{}{}
		}

	}
}

func (w *Writer) syncSync() {
	for {
		select {
		case f, ok := <-w.syncFileBuf:
			if !ok {
				return
			}
			f.Sync()
			f.Close()
		}
	}
}
