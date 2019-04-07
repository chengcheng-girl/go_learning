package sync

import (
	"context"
	"testing"
	"time"
)

func isCancelContext(ctx context.Context) bool  {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}


func TestCancelContext(t *testing.T)  {
	ctx,cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int,ctx context.Context) {
			for {
				if isCancelContext(ctx) {
					break
				}
				time.Sleep(time.Microsecond*10)
			}
			t.Log(i,"canceled")
		}(i,ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
