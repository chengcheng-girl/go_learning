package sync

import (
	"fmt"
	"github.com/wonderivan/logger"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	logger.Info("int:", i)
	return fmt.Sprintf("This result is %d", i)
}

func firstResponse() string {
	number := 10
	ch := make(chan string, number)
	for i := 0; i < number; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

//凡是有结果就返回输出
func TestFirstResponse(t *testing.T) {
	logger.Warn("before:", runtime.NumGoroutine())
	logger.Warn("result:", firstResponse())
	time.Sleep(time.Second * 1)
	logger.Warn("after:", runtime.NumGoroutine())

}
