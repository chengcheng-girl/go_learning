package obj_pool

import (
	"errors"
	"time"
)

type Object struct {
}

type ObjectPool struct {
	chanBuffer chan *Object
}

func newObject(number int) *ObjectPool {
	objPool := &ObjectPool{}
	objPool.chanBuffer = make(chan *Object, number)
	for i := 0; i < number; i++ {
		objPool.chanBuffer <- &Object{}
	}
	return objPool
}

func (pool *ObjectPool) getObject(timeout time.Duration) (*Object, error) {
	select {
	case ret := <-pool.chanBuffer:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (pool *ObjectPool) updateObject(obj *Object) error {
	select {
	case pool.chanBuffer <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
