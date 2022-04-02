package syncx

import (
	"sync"
	"time"
)

// default refresh time.
const defaultRefreshInterval = time.Second

// 资源操作
type (
	ConstantResourcesOpt func(resources *ConstantResources)
	ConstantFetch        func() (interface{}, error)
	ConstantResources    struct {
		fetch           ConstantFetch   // 刷新资源方法
		resources       interface{}     // 资源
		refreshInterval time.Duration   // 刷新间隔
		lastTime        *AtomicDuration // 最新的时间
		err             error           // 错误
		lock            sync.RWMutex    // 读写互斥锁
	}
)

func NewConstantResources(fn ConstantFetch,
	opts ...ConstantResourcesOpt) *ConstantResources {

	constant := ConstantResources{
		fetch:           fn,
		refreshInterval: defaultRefreshInterval,
		lastTime:        NewAtomicDuration(),
	}

	for _, opt := range opts {
		opt(&constant)
	}

	return &constant
}

// get resources
func (con *ConstantResources) Take() (interface{}, error) {
	con.lock.RLock()
	resource := con.resources
	con.lock.RUnlock()

	if resource != nil {
		return resource, nil
	}

	con.attemptRefresh(func() {
		resource, err := con.fetch()
		con.lock.Lock()
		if err != nil {
			con.err = err
		} else {
			con.resources = resource
		}
		con.lock.Unlock()
	})

	con.lock.RLock()
	resource, err := con.resources, con.err
	con.lock.RUnlock()

	if err != nil {
		return nil, err
	}

	return resource, err
}

// attempt refresh resources.
func (con *ConstantResources) attemptRefresh(callback func()) {
	now := time.Duration(time.Now().Unix())
	last := con.lastTime.Load()

	if last == 0 || last+con.refreshInterval < now {
		// refresh resources
		con.lastTime.Set(now)
		callback()
	}
}

// set resources refresh interval time.
func SetRefreshInterval(t time.Duration) ConstantResourcesOpt {
	return func(ins *ConstantResources) {
		ins.refreshInterval = t
	}
}
