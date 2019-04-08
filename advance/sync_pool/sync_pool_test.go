package sync_pool

import (
	"sync"
	"testing"
)

//object cache

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 100
		},
	}
	pool.Put(1000)
	pool.Put(666)
	v := pool.Get().(int)
	t.Log("before value:", v)
	//runtime.GC() //gc会清空sync.pool中的对象
	v = pool.Get().(int)
	t.Log("after value:", v)
	v = pool.Get().(int)
	t.Log("after value:", v)
	//fifo queue
}

func TestSyncMultiPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 9999
		},
	}
	pool.Put(100)
	pool.Put(200)
	pool.Put(300)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			t.Log("result:", pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}

//通过复用　降低复杂对象的创建和gc的代价
//协成安全　会有锁的开销
//生命周期受gc影响
