package test

import (
	"testing"
	"time"

	"github.com/372572571/xq/syncx"
)

func TestDoneChan(t *testing.T) {
	var fibTests = []struct {
		end time.Duration // 发送信号时间
	}{
		{3},
		{0},
		{1},
		{1},
		{1},
	}

	for _, tt := range fibTests {
		ins := syncx.NewDoneChan()
		timer := time.NewTimer(tt.end * time.Second)

		go func() {
			<-timer.C
			ins.Close()
		}()
		begin := time.Now()
		<-ins.Done()
		t.Logf("time consuming:%d millisecond  \n", time.Now().Sub(begin)/time.Millisecond)
	}
}
