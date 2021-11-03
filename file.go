package zapper

import (
	"strings"
	"time"
)

func GenLogFilepath(filepath string, nextTheTime time.Time) string {
	fileSlice := strings.Split(filepath, ".")
	return strings.Join([]string{fileSlice[0], nextTheTime.Format("2006-01-02 15:04:05"), ".log"}, "")
}


