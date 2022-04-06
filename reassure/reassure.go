package reassure

import (
	"fmt"
)

// 捕获
func Reassure(fns ...func()) {
	for _, fn := range fns {
		fn()
	}

	if err := recover(); err != nil {
		fmt.Println("[error]", err)
	}
}
