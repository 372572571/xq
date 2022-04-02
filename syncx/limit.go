package syncx

import (
	"errors"

	"github.com/372572571/xq/lang"
)

var LimitedErr = errors.New("limitation in..")

// limit bucket
type Limit struct {
	pool chan lang.PlaceholderType
}

func NewLimit(n int) *Limit {
	return &Limit{
		pool: make(chan lang.PlaceholderType, n),
	}
}

// if pool full wait..
func (l Limit) PutIn() {
	l.pool <- lang.Placeholder
}

// If full return false or not full return true
// not wait
func (l Limit) TryPut() bool {
	select {
	case l.pool <- lang.Placeholder:
		return true
	default:
		return false
	}
}

// take one
func (l Limit) Take() error {
	select {
	case <-l.pool:
		return nil
	default:
		return LimitedErr
	}
}
