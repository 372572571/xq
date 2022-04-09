package reassure

import "log"

// 捕获
func Reassure(fns ...func()) {
	for _, fn := range fns {
		fn()
	}

	if err := recover(); err != nil {
		log.Default().Println(err)
	}
}
