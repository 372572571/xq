package main

import (
	"fmt"
	"time"

	"github.com/372572571/xq/database"
	"github.com/372572571/xq/logx"
	"github.com/372572571/xq/syncx"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	sql()
	return
	// logtest()
	timetest()
	return
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "test"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func timetest() {
	fmt.Println(-60 * 24 * time.Hour)
	fmt.Println(time.Unix(1650549931, 0))
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

func testsql() {
	var db = &database.Database{}
	test_db := db.Unsafe().Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("test = 1")
	})
	test_db = test_db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("test2 = 2")
	})

	// test_db.
}

func sql() {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "123456",
		"172.16.13.4", "3306", "bang")
	db, err := gorm.Open(mysql.Open(connArgs), &gorm.Config{})

	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}
	// var

	// join pay_order on task.task_no=pay_order.order_no and pay_order.status = ?
	var data []struct {
		id int64
	}
	var count int64
	var res = db.Debug().Exec(` UPDATE task  
	SET task.status = ?
	where task.id in ? and task.status = ?`,
		41,                            // 要更新到的状态
		[]int64{87198277894150, 2222}, // 任务id
		40,
	)
	logx.Info(count)
	logx.Info(data)
	logx.Info(res.RowsAffected)
	// logx.Info(res.Columns())
	// stmt := db.Session(&gorm.Session{DryRun: true}).Table("user").Where("id = 1").Find(&struct{}{}).Statement
	// // stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	// // stmt.Vars         //=> []interface{}{1}
	// // logx.Infof("%s", stmt.SQL.String())
	// var data []struct{
	// 	Name string
	// }
	// db.Debug().Raw(stmt.SQL.String()).Limit(0).Find(&data)

	// logx.Info(data)
}
