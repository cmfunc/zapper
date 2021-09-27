package zapper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCron(t *testing.T) {
	var filepath string = "./test_cron_rotate.log"

	syncCycle := time.Hour
	wr := NewFileWriter(filepath, syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)

	syncLock := make(chan struct{}, 0)
	go func() {
		for i := 0; i <= 180; i++ {

			logger.Error("test cron zaper ",
				zap.Int("int", 10),
				zap.Error(errors.New("text string")),
				zap.String("key string", "val string"),
				zap.Time("timr", time.Now()),
			)

			time.Sleep(time.Second * 1)
		}

		time.Sleep(time.Second * 20)
		logger.Sync()
		syncLock <- struct{}{}
	}()

	<-syncLock

}
