package threading

import (
	"bytes"
	"runtime"
	"strconv"

	"github.com/372572571/xq/reassure"
)

// RoutineId is only for debug, never use it in production.
func RoutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

func GoSafe(fn func()) {
	go RunSafe(fn)
}

// 捕获异常
func RunSafe(fn func()) {
	defer reassure.Reassure()
	fn()
}
