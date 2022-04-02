package syncx

import "sync"

type (

	// 使用key来分组支持顺序调用
	// key = "A"  func A1 -> func A2
	// key = "B"  func A2 -> func A1
	// 组b的调用不会受到A组的影响
	LockedCalls interface {
		Do(key string, fn func() (interface{}, error)) (interface{}, error)
	}

	LockedGroup struct {
		lock  sync.Mutex
		group map[string]*sync.WaitGroup // 分组锁
	}
)

func NewLockedGroup() LockedCalls {
	return &LockedGroup{
		group: make(map[string]*sync.WaitGroup),
	}
}

func (lo *LockedGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {

begin:
	lo.lock.Lock()
	// try get group lock.
	if wg, ok := lo.group[key]; ok {
		// 如果有则解锁并等待
		lo.lock.Unlock()
		wg.Wait()
		goto begin // end of previous func.  try to continue take lock.
	}

	// if not key then exec self function
	return lo.makeCalls(key, fn)
}

func (lo *LockedGroup) makeCalls(key string, fn func() (interface{}, error)) (interface{}, error) {
	// create new lock
	var wg sync.WaitGroup
	wg.Add(1)
	lo.group[key] = &wg
	lo.lock.Unlock()

	// clear lock and done
	defer func() {
		lo.lock.Lock()
		delete(lo.group, key)
		lo.lock.Unlock()
		wg.Done()
	}()

	// exec
	return fn()
}
