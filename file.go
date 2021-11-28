package zapper

import (
	"strings"
	"time"
)

// GenLogFilepath 生成带整点时间标记的日志文件名
// filepath 文件默认地址
// nextTheTime下一个整点时间
// 返回整点时间日志文件地址
func GenLogFilepath(filepath string, nextTheTime time.Time, tk time.Duration) string {
	fileSlice := strings.Split(filepath, ".")
	var formatStyle string
	switch tk {
	case time.Nanosecond:
		formatStyle = "20060102150405000000000"
	case time.Second:
		formatStyle = "20060102150405"
	case time.Minute:
		formatStyle = "200601021504"
	case time.Hour:
		formatStyle = "2006010215"
	case TheDay:
		formatStyle = "20060102"
	case TheMonth:
		formatStyle = "200601"
	case TheYear:
		formatStyle = "2006"
	default:
		formatStyle = "2006010215"
	}
	return strings.Join([]string{fileSlice[0], nextTheTime.Format(formatStyle), ".log"}, "")
}
