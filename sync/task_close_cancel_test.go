package sync

import (
	"testing"
	"time"
)

func isCancel(cancel chan struct{}) bool  {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

func cancelSign(cancel chan struct{})  {
	cancel<- struct{}{}
}
func cancelChan(cancel chan struct{})  {
	close(cancel)
}

func TestCancel(t *testing.T)  {
	cancelCh := make(chan struct{},1)
	for i:=0;i<5;i++ {
		go func(i int,cancelCh chan struct{}) {
			for {
				if isCancel(cancelCh) {
					break
				}
				time.Sleep(time.Microsecond*10)
			}
			t.Log(i,"canceled")
		}(i,cancelCh)
	}
	cancelChan(cancelCh)
	time.Sleep(time.Second * 1)
}
