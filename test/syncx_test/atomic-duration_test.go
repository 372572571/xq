package test

import (
	"testing"
	"time"

	"github.com/372572571/xq/syncx"
)

func TestAtomicDuration(t *testing.T) {
	var fibTests = []struct {
		in  time.Duration // 运行时间
		tow time.Duration
	}{
		{9, 9},
		{10, 1},
		{1, 2},
		{0, 0},
		{1, 1},
	}

	for _, tt := range fibTests {
		ins := syncx.NewAtomicDuration()
		ins.ForAtomicDuration(tt.in)

		b := ins.CompareAndSwap(tt.in, tt.tow)

		t.Logf("in:%d  tow:%d  bool %t \n", tt.in, tt.tow, b)
	}
}

func TestAtomicDurationLoad(t *testing.T) {
	var fibTests = []struct {
		in  time.Duration
		tow time.Duration
	}{
		{9, 9},  // false
		{10, 1}, // true
		{1, 2},  // false
		{0, 0},  // false
		{1, 1},  // false
	}

	for _, tt := range fibTests {
		ins := syncx.NewAtomicDuration()
		ins.Set(10) // 10 true
		// tt.in = 10 时才会是true  数值数值成功互换
		c := ins.CompareAndSwap(tt.in, tt.tow)
		t.Logf("in:%d  tow:%d  bool %t value %d \n", tt.in, tt.tow, c, ins.Load())
	}
}
