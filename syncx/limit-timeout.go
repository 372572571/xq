package syncx

import (
	"errors"
	"time"
)

var LimitTimeOutError = errors.New("timeout put")

type (
	LimitTimeOut struct {
		limit *Limit
		cond  *Cond
	}
)

func NewLimitTimeOut(limit_n int) *LimitTimeOut {
	return &LimitTimeOut{
		limit: NewLimit(limit_n),
		cond:  NewCond(),
	}
}

func (l *LimitTimeOut) Put(timeout time.Duration) error {
	if l.limit.TryPut() { // 如果能够放进去直接返回
		return nil
	}

	var ok bool

	for {
		// 如果不能直接放进去,使用带有超时的阻塞的方法.
		// 阻塞结束时再次尝试放入
		// 阻塞结束的方式超时
		// 主动和调用l.cond.singe
		timeout, ok = l.cond.WaitWithTimeout(timeout)
		if ok && l.limit.TryPut() {
			return nil
		}

		if timeout <= 0 {
			return LimitTimeOutError
		}
	}
}

func (l *LimitTimeOut) Take() error {

	if err := l.limit.Take(); err != nil {
		return err
	}
	// 主动取出后,尝试通知,可能存在的存入阻塞.继续运行.
	l.cond.Signal()
	return nil
}
