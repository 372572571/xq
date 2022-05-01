package redisx

import "github.com/gomodule/redigo/redis"

type (
	Values interface {
		String(reply interface{}, err error) (string, error)
	}
	values struct{}
)

func (v values) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}
