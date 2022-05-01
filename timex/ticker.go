package timex

import "time"

type (
	Ticker interface {
		Chan() <-chan time.Time
		Stop()
	}

	realTicker struct {
		*time.Ticker
	}
)

func NewTicker(t time.Duration) Ticker {
	return &realTicker{time.NewTicker(t)}
}

func (this *realTicker) Chan() <-chan time.Time {
	return this.C
}
