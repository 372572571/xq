package single

import "github.com/hibiken/asynq"

type IAsynqc interface {
	Client() *asynq.Client
}

type asynqc struct {
	client *asynq.Client
	opt    *optsions
}

func NewClient(funcs ...OptFunc) IAsynqc {
	var c = &asynqc{}
	c.opt = &optsions{redis: asynq.RedisClientOpt{}}

	for _, f := range funcs {
		f(*c.opt)
	}

	c.client = asynq.NewClient(c.opt.redis)

	return c
}

func (c *asynqc) Client() *asynq.Client {
	return c.client
}
