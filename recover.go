// zaper
// 捕获Panic 避免日志导致程序退出
package zaper

import "go.uber.org/zap"

func logcatch()  {
	if err := recover(); err!= nil{
		logger.Error("Error: ", zap.Any("error",err))
	}
}