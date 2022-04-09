package main

import (
	"time"

	"github.com/372572571/xq/logx"
	"github.com/372572571/xq/syncx"
	"go.uber.org/zap"
)

// func main() {
// 	t := &time.Time{}
// 	t2 := &time.Time{}
// 	fmt.Println(t)
// 	fmt.Printf("%p \n",t)
// 	fmt.Printf("%p \n",t2)
// 	fmt.Println((*time.Time)(t))
// }

// func main() {
// 	data := make(chan bool, 2)
// 	data <- bool(true)
// 	fmt.Println(1)
// 	data <- bool(true)
// 	fmt.Println(2)
// 	go func() {
// 		t := time.NewTicker(time.Second * 10)
// 		<-t.C
// 		<-data
// 		fmt.Println("time over")
// 	}()
// 	data <- bool(true)
// 	fmt.Println(3)
// 	fmt.Println("over")
// }

type data struct {
	Value int64
}

func main() {
	// logtest()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "test"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

var log *logx.Logger

func logtest() {
	defer func() {
		if err := recover(); err != nil {
			log.DPanic(err)
		}
	}()
	log = logx.NewLog(logx.Config{Info: "./logs/doss.info", Error: "logs/doss.error", Panic: "logs/doss.panic"})
	log.Errorf("test error")

	log.Warnf("test warnf")

	log.Infof("test infof")

	cond := syncx.NewCond()
	cond.WaitWithTimeout(time.Second * 2)
	panic("error")

	// var data []int
	// fmt.Println(data[3])
	// for {
	// }

}
