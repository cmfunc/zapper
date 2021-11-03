package zapper

import (
	"time"
)

const (
	TheNanosecond time.Duration = time.Nanosecond
	TheSecond                   = time.Second
	TheMinute                   = time.Minute
	TheHour                     = time.Hour
	TheDay                      = time.Hour * 24
	TheMonth                    = time.Hour * 720
	TheYear                     = time.Hour * 8760
)

// NextTheTime 根据tk指定的单位，确定t的整点时间
func NextTheTime(t time.Time, tk time.Duration) time.Time {
	t = t.Add(tk)
	switch tk {
	case time.Nanosecond:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	case time.Second:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	case time.Minute:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	case time.Hour:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	case TheDay:
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case TheMonth:
		t = time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
	case TheYear:
		t = time.Date(t.Year(), 0, 0, 0, 0, 0, 0, t.Location())
	default:
		t = t.Add(-tk)
		t = t.Add(time.Hour)
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	}
	return t
}
