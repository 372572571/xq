package syncx

import (
	"fmt"
	"sync"
	"time"
)

type (
	PoolOption func(*Pool)

	node struct {
		last      time.Duration // put last time
		resources interface{}
		next      *node
	}

	Pool struct {
		limit   int           // limit number
		created int           // already create number
		max     time.Duration // node retain time
		lock    sync.Locker
		cond    *sync.Cond
		head    *node
		create  func() interface{}
		destroy func(interface{})
	}
)

func NewPool(n int, create func() interface{}, destroy func(interface{}), opts ...PoolOption) (*Pool, error) {
	if n <= 0 {
		return nil, fmt.Errorf("cannot limit zero.")
	}

	lock := new(sync.Mutex)
	pool := Pool{
		limit:   n,
		lock:    lock,
		cond:    sync.NewCond(lock),
		create:  create,
		destroy: destroy,
	}

	for _, opt := range opts {
		opt(&pool)
	}

	return &pool, nil
}

func (p *Pool) Take() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()

	for {

		// if head
		if p.head != nil {
			head := p.head
			p.head = head.next
			// retain time inspect
			if p.max > 0 && head.last+p.max < time.Duration(time.Now().UnixNano()) {
				p.created -= 1
				p.destroy(head.resources)
				continue
			} else {
				return head.resources
			}
		}

		// if limit
		if p.created < p.limit {
			p.created += 1
			return p.create()
		}

		// wait put resources
		p.cond.Wait()
	}

}

// put resources
func (p *Pool) Put(value interface{}) {
	if value == nil {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()
	p.head = &node{
		resources: value,
		last:      time.Duration(time.Now().Nanosecond()),
		next:      p.head,
	}

	// send done sign
	p.cond.Signal()
}

// set resources retain max time
func SetMaxTime(value time.Duration) PoolOption {
	return func(p *Pool) {
		p.max = value
	}
}
