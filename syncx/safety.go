package syncx

import "sync"

// 保证安全调用
type Safety struct {
	lock sync.Mutex // 互斥锁
}

// 互斥锁保障func
func (one *Safety) Guarantee(fun func()) {
	employee(&one.lock, fun)
}

// A Locker represents an object that can be locked and unlocked.
func employee(lock sync.Locker, fun func()) {
	lock.Lock()
	defer lock.Unlock()
	fun()
}
