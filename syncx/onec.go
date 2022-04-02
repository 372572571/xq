package syncx

import "sync"

// exec once func
func OnceFunc(fn func()) func() {
	once := new(sync.Once)
	return func() {
		once.Do(fn)
	}
}
