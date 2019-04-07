package sync

import (
	"github.com/wonderivan/logger"
	"testing"
	"time"
)

func TestCsp(t *testing.T)  {
	retCh :=syncService()
	otherService()
	//get data from channel
	t.Log(<-retCh)

}
func TestCspSelect(t *testing.T)  {
	select {
	case retCh :=<-syncService():
		t.Log(retCh)
	case <-time.After(time.Second*11):
		t.Error("time out")
	}

	otherService()
	//get data from channel


}

func otherService() {
	time.Sleep(time.Second * 5)
}

func service() string  {
	time.Sleep(time.Second * 10)
	return "Done"
}

func syncService() chan string {
	reCh :=make(chan string) //单一模式　阻塞
	//reCh :=make(chan string,1)//缓存模式
	go func() {
		ret:=service()
		logger.Info("ret:",ret)
		//insert data to channel
		reCh <- ret
		logger.Info("end")
	}()
	return reCh
}
