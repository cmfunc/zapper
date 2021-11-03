package zapper

import (
	"testing"
	"time"
)

func TestNextTheTime(t *testing.T) {
	result := time.Date(2021, 10, 9, 1, 0, 0, 0, time.Local)
	ti := time.Date(2021, 10, 9, 0, 23, 23, 45, time.Local)
	tN := NextTheTime(ti, time.Hour)
	if !tN.Equal(result) {
		t.Errorf("生成 next_time 不正确 result:%s nextTheTime:%s", result, tN)
		return
	}
	
	result = time.Date(2021, 10, 9, 1, 1, 0, 0, time.Local)
	ti = time.Date(2021, 10, 9, 1, 0, 23, 45, time.Local)
	tN = NextTheTime(ti, time.Minute)
	if !tN.Equal(result) {
		t.Errorf("生成 next_time 不正确 result:%s nextTheTime:%s", result, tN)
		return
	}

}
