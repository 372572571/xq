package single

import (
	"fmt"

	"github.com/hibiken/asynq"
)

type IAsynqc interface {
	Client() *asynq.Client
}

type asynqc struct {
	client *asynq.Client
	opt    *optsions
}

func NewClient(config Config, funcs ...OptFunc) IAsynqc {
	var c = &asynqc{}
	c.opt = &optsions{redis: asynq.RedisClientOpt{}}

	c.opt.redis.Addr     = config.Addr
	c.opt.redis.DB       = config.Index
	c.opt.redis.Password = config.Password

	for _, f := range funcs {
		f(*c.opt)
	}

	fmt.Println(c.opt.redis)

	c.client = asynq.NewClient(c.opt.redis)

	return c
}

func (c *asynqc) Client() *asynq.Client {
	return c.client
}
