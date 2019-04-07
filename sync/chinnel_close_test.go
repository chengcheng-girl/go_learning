package sync

import (
	"github.com/wonderivan/logger"
	"sync"
	"testing"
)

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch :=make(chan int,1)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}

func TestChannelCloseDefault(t *testing.T) {
	var wg sync.WaitGroup
	ch :=make(chan int,1)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiverDefault(ch,&wg)
	wg.Wait()
}

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				logger.Info(data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func dataReceiverDefault(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			data := <-ch
			logger.Info(data)
		}
		wg.Done()
	}()

}
