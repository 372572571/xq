package alone

import (
	"github.com/gomodule/redigo/redis"
)

type mode struct {
	pool *redis.Pool
}

func (sm *mode) GetPool() *redis.Pool {
	return sm.pool
}

func (sm *mode) GetConn() redis.Conn {
	return sm.pool.Get()
}

func (sm *mode) NewConn() (redis.Conn, error) {
	return sm.pool.Dial()
}

func (sm *mode) Close() error {
	return sm.pool.Close()
}

func New(fns ...OptFunc) *mode {
	opts := options{}
	for _, fn := range fns {
		fn(&opts)
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", opts.addr, opts.dialOpts...)
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(pool)
	}
	return &mode{pool}
}
