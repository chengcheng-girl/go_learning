package sync

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T){
	count :=0
	for i:=0;i<5000;i++{
		go func() {
			count++
		}()
	}
	time.Sleep(1* time.Second)
	t.Log("count num:",count)
}


func TestCounterSafe(t *testing.T){
	var mux sync.Mutex
	count :=0
	for i:=0;i<5000;i++{
		go func() {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			count++
		}()
	}
	// sync running is so fast.
	time.Sleep(1* time.Second)
	t.Log("count num:",count)
}



func TestCounterWaitGroup(t *testing.T){
	var mux sync.Mutex
	var wg sync.WaitGroup
	count :=0
	for i:=0;i<5000;i++{
		wg.Add(1)
		go func() {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("count num:",count)
}
