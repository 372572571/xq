package single

import (
	"github.com/hibiken/asynq"
)

type IAsynqSingle interface {
	NewServeMux() *asynq.ServeMux
	Server() *asynq.Server
}

type asynqs struct {
	server *asynq.Server
	opt    *optsions
}

func NewAsynqxServer(config Config, opts ...OptFunc) IAsynqSingle {
	var opt = &optsions{
		redis:  asynq.RedisClientOpt{},
		config: asynq.Config{},
	}

	opt.redis.Addr     = config.Addr
	opt.redis.Password = config.Password
	opt.redis.DB       = config.Index
	
	for _, v := range opts {
		v(*opt)
	}

	return &asynqs{
		server: asynq.NewServer(opt.redis, opt.config),
		opt:    opt,
	}
}

func (x *asynqs) Server() *asynq.Server {
	return x.server
}

func (x *asynqs) NewServeMux() *asynq.ServeMux {
	return asynq.NewServeMux()
}
