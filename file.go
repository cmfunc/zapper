package zapper

import (
	"strings"
	"time"
)

// GenLogFilepath 生成带整点时间标记的日志文件名
// filepath 文件默认地址
// nextTheTime下一个整点时间
// 返回整点时间日志文件地址
func GenLogFilepath(filepath string, nextTheTime time.Time) string {
	fileSlice := strings.Split(filepath, ".")
	return strings.Join([]string{fileSlice[0], nextTheTime.Format("2006-01-02 15:04:05"), ".log"}, "")
}


