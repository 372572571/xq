package single

import "github.com/hibiken/asynq"

type optsions struct {
	redis  asynq.RedisClientOpt
	config asynq.Config
}

type OptFunc = func(opt optsions)

// redis 地址
func Addr(value string) OptFunc {
	return func(opt optsions) {
		opt.redis.Addr = value
	}
}

// redis 密码
func Pwd(value string) OptFunc {
	return func(opt optsions) {
		opt.redis.Password = value
	}
}

// 并发数量
func Concurrency(value int) OptFunc {
	return func(opt optsions) {
		opt.config.Concurrency = value
	}
}

// 队列配置
func Queues(value map[string]int) OptFunc {
	return func(opt optsions) {
		opt.config.Queues = value
	}
}

// 严格按照优先级执行队列
func StrictQueue(value bool) OptFunc {
	return func(opt optsions) {
		opt.config.StrictPriority = value
	}
}

// 默认配置
func DefaultConfig(addr string, pwd string) []OptFunc {
	var funcs = []OptFunc{}

	funcs = append(funcs, func(opt optsions) {
		opt.redis.Addr = addr
		opt.redis.Password = pwd
		opt.config.Concurrency = 5
		opt.config.Queues = map[string]int{
			"up":      3,
			"default": 2,
			"low":     1,
		}
	})

	return funcs
}
