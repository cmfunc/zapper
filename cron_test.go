package zaper

import (
	"errors"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNewFileWriter(t *testing.T) {
	filepath := "./TestNewFileWriter.log"
	w :=  NewFileWriter(filepath).Load()
	
	w.(*os.File).WriteString("shgkksk nj")
	defer w.(*os.File).Close()
}

func TestCron(t *testing.T) {
	var filepath string = "./test_cron_rotate.log"

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", filepath)

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
