package zaper

import (
	"errors"
	"testing"
	"time"

	"zaper/gocommons-develop-logging/logging"
)

func TestLoggingInfo(t *testing.T) {
	msg := map[string]string{"AAAA": "BBBBB"}
	c := logging.LogConfig{Path: "./", File: "logging_benchmark.log", Mode: 1, Rotate: false, Debug: false}
	logger := logging.NewLogger(&c)
	header := logging.LogHeader{LogId: "abc"}

	logger.Info(&header, msg)

	time.Sleep(100 * time.Millisecond)
}

func BenchmarkLoggingInfo(b *testing.B) {
	msg := map[string]string{"AAAA": "BBBBB"}
	c := logging.LogConfig{Path: "./", File: "logging_benchmark.log", Mode: 1, Rotate: false, Debug: false}
	logger := logging.NewLogger(&c)
	header := logging.LogHeader{LogId: "abc"}

	for i := 0; i < b.N; i++ {
		logger.Error(&header, msg, 123, "234", []interface{}{123, "32"}, logger, errors.New("text string"))
	}

	time.Sleep(100 * time.Millisecond)

}
