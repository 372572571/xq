package threading

type WorkerGroup struct {
	fn      func() // 执行的方法
	workers int64  // 工作go携程数量
}

func NewWorkerGroup(fn func(), worker_number int64) *WorkerGroup {
	return &WorkerGroup{fn: fn, workers: worker_number}
}

func (worker WorkerGroup) Start() {
	group := NewRoutineGroup()
	for i := int64(0); i < worker.workers; i++ {
		group.RunSafe(worker.fn)
	}
	group.wg.Wait()
}
