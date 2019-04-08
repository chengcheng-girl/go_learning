package sync

import (
	"github.com/wonderivan/logger"
	"runtime"
	"testing"
	"time"
)

func allResponse() string {
	number := 10
	ch := make(chan string, number)
	for i := 0; i < number; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalReturn := ""
	for i := 0; i < number; i++ {
		finalReturn += <-ch + "\n"
	}
	return finalReturn
}

//所有任务执行完成
func TestAllResponse(t *testing.T) {
	logger.Warn("before:", runtime.NumGoroutine())
	logger.Warn("result:", allResponse())
	time.Sleep(time.Second * 1)
	logger.Warn("after:", runtime.NumGoroutine())

}
