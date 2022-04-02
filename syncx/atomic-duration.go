package syncx

import (
	"sync/atomic"
	"time"
)

type AtomicDuration int64

func NewAtomicDuration() *AtomicDuration {
	return new(AtomicDuration)
}

func (duration *AtomicDuration) Set(v time.Duration) {
	// 把v储存到 duration (*int64 指针) 地址中
	atomic.StoreInt64((*int64)(duration), int64(v))
}

func (duration *AtomicDuration) Load() time.Duration {
	// 从int64的指针中读取数值转换为 time.Duration
	return time.Duration(atomic.LoadInt64((*int64)(duration)))
}

func (duration *AtomicDuration) ForAtomicDuration(v time.Duration) *AtomicDuration {
	d := NewAtomicDuration()
	d.Set(v)
	return d
}

// compare old val swap
func (duration *AtomicDuration) CompareAndSwap(old, val time.Duration) bool {
	return atomic.CompareAndSwapInt64((*int64)(duration), int64(old), int64(val))
}
