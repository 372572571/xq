package syncx

import "sync"

type (
	SingleCall interface {
		Do(key string, fn func() (interface{}, error)) (interface{}, error)
		DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
	}

	call struct {
		wg        sync.WaitGroup
		resources interface{}
		err       error
	}

	singleCall struct {
		lock  sync.Mutex
		calls map[string]*call
	}
)

func NewSingleCall() SingleCall {
	return &singleCall{
		calls: map[string]*call{},
	}
}

func (s *singleCall) createCall(key string) (*call, bool) {
	s.lock.Lock()
	var done bool = true
	if c, ok := s.calls[key]; ok {
		s.lock.Unlock()
		c.wg.Wait()    // 说明有相同key的回调在调用此时阻塞
		return c, done // 对应key的回调完成了
	}

	// 没有相同的key调用(自己来)
	call := new(call)
	call.wg.Add(1)
	s.calls[key] = call
	s.lock.Unlock()

	return call, !done
}

func (s *singleCall) makeCall(c *call, key string, fn func() (interface{}, error)) {

	// 释放资源并通知其他等待进程解除阻塞获取资源
	defer func() {
		s.lock.Lock()
		delete(s.calls, key)
		s.lock.Unlock()
		c.wg.Done()
	}()

	// 调用回调获取资源
	c.resources, c.err = fn()
}

func (s *singleCall) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	call, done := s.createCall(key)

	if done {
		return call.resources, call.err
	}

	s.makeCall(call, key, fn)

	return call.resources, call.err

}

func (s *singleCall) DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error) {
	call, done := s.createCall(key)

	if done {
		return call.resources, !done, call.err
	}

	s.makeCall(call, key, fn)

	return call.resources, done, call.err

}
