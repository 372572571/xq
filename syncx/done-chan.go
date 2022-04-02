package syncx

import (
	"sync"

	"github.com/372572571/xq/lang"
)

type DoneChan struct {
	done chan lang.PlaceholderType
	one  sync.Once
}

func NewDoneChan() *DoneChan {
	return &DoneChan{
		done: make(chan lang.PlaceholderType),
	}
}

// close once done chan
func (dc *DoneChan) Close() {
	dc.one.Do(func() {
		close(dc.done)
	})
}

func (dc *DoneChan) Done() chan lang.PlaceholderType {
	return dc.done
}
