package threading

import (
	"github.com/372572571/xq/lang"
	"github.com/372572571/xq/reassure"
)

type TaskRunner struct {
	limitChan chan lang.PlaceholderType
}

// concurrent  can concurrent number
func NewTaskRunner(concurrent int) *TaskRunner {
	return &TaskRunner{
		limitChan: make(chan lang.PlaceholderType, concurrent),
	}
}

// concurrent call func
func (tr *TaskRunner) Schedule(fn func()) {
	tr.limitChan <- lang.Placeholder
	go func() {
		defer reassure.Reassure(func() {
			<-tr.limitChan
		})

		fn()
	}()
}
