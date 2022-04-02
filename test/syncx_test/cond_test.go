package test

import (
	"testing"
	"time"

	"github.com/372572571/xq/syncx"
)

func TestCond(t *testing.T) {
	var fibTests = []struct {
		in  time.Duration // 运行时间
		end time.Duration // 发送信号时间
	}{
		{1, 9},
		{2, 10},
		{1, 1},
		{1, 0},
		{2, 1},
	}

	for _, tt := range fibTests {
		ins := syncx.NewCond()
		timer := time.NewTimer(tt.end * time.Second)

		go func() {
			<-timer.C
			ins.Signal()
		}()

		res, b := ins.WaitWithTimeout(tt.in * time.Second)

		t.Logf("time surplus:%d  timeout bool %t \n", res, b)
	}
}
