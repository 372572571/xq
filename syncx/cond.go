package syncx

import (
	"time"

	"github.com/372572571/xq/lang"
)

// 阻塞到等待通知
type Cond struct {
	signal chan lang.PlaceholderType
}

func NewCond() *Cond {
	return &Cond{
		signal: make(chan lang.PlaceholderType),
	}
}

// false : timer retrun
// true  : <-c.signal
func (c *Cond) WaitWithTimeout(timeout time.Duration) (time.Duration, bool) {
	timer := time.NewTicker(timeout)
	defer timer.Stop()

	begin := time.Now()

	select {
	case <-c.signal:
		end := time.Now().Sub(begin)
		return timeout - end, true
	case <-timer.C:
		return 0, false
	}
}

// wait signal
func (c *Cond) Wait() {
	<-c.signal
}

// send signal
func (c *Cond) Signal() {
	select {
	case c.signal <- lang.Placeholder:
	default:
	}
}
