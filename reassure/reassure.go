package reassure

import (
	"github.com/372572571/xq/logx"
)

// 捕获
func Reassure(fns ...func()) {
	for _, fn := range fns {
		fn()
	}

	if err := recover(); err != nil {
		logx.Error(err)
	}
}
