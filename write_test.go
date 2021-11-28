package zapper

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewWriter(t *testing.T) {
	var filepath string = "/Users/yymt/Documents/zaper/test_cron_rotate.log"

	syncCycle := time.Minute
	wr := NewWriter(filepath, syncCycle, 1024)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	SetDefaultLogger(logger)

	syncLock := make(chan struct{}, 0)
	go func() {
		for i := 0; i <= 240; i++ {
			err:=errors.New("text string001")
			Error("test cron zaper ",
				zap.Int("int", 10),
				zap.Error(fmt.Errorf("text string: %w",err)),
				zap.String("msg", "val string1"),
				zap.String("msg", "val string2"),
				zap.String("msg", "val string3"),
				zap.Time("time", time.Now()),
			)

			time.Sleep(time.Second * 1)
		}

		time.Sleep(time.Second * 20)
		Sync()
		syncLock <- struct{}{}
	}()

	<-syncLock

}

func TestAbs(t *testing.T)  {
	s:="2021-10-09 15:00:00.log"
	if !filepath.IsAbs(s){
		f,e:=filepath.Abs(s)
		println(f,e)
	}

	f:="/Users/yymt/Documents/zaper/zapper2021-11-08 17:00:00.log"
	f2,e:=filepath.Abs(f)
	t.Log(f2,e)
}
