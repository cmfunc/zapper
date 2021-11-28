package zapper

import (
	"testing"
	"time"
)

func TestGenLogFilepath(t *testing.T) {
	result := "my2021-10-09 00:00:00.log"
	file := "my.log"
	dur := time.Date(2021, 10, 9, 0, 0, 0, 0, time.Local)
	log_name := GenLogFilepath(file, dur,time.Minute)
	if log_name != result {
		t.Errorf("生成log_name不正确 result:%s log_naem:%s", result, log_name)
		return
	}
	t.Log("Success")
}
