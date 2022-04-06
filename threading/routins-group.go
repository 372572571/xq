package threading

import "sync"

// 正常顺序执行
type RoutineGroup struct {
	wg sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

// 不捕获异常
func (r *RoutineGroup) Run(fn func()) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		fn()
	}()
}

// 捕获异常方式
func (r *RoutineGroup) RunSafe(fn func()) {
	r.wg.Add(1)
	GoSafe(func() {
		defer r.wg.Done()
		fn()
	})
}

// 等待完成
func (r *RoutineGroup) Wait() {
	r.wg.Wait()
}
