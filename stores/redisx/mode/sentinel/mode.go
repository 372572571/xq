package sentinel

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
)

type mode struct {
	st   *Sentinel
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

func New(set_fun ...OptFunc) *mode {
	opts := options{
		// addrs: []string{
		// 	"127.0.0.1:26379",
		// },
		// masterName: "master",
	}
	for _, optFunc := range set_fun {
		optFunc(&opts)
	}
	if len(opts.sentinelDialOpts) == 0 {
		opts.sentinelDialOpts = opts.dialOpts
	}
	st := &Sentinel{
		Addrs:      opts.addrs,
		MasterName: opts.masterName,
		Pool: func(addr string) *redis.Pool {
			return &redis.Pool{
				Wait:    true,
				MaxIdle: runtime.GOMAXPROCS(0),
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", addr, opts.sentinelDialOpts...)
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) (err error) {
					_, err = c.Do("PING")
					return
				},
			}
		},
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			addr, err := st.MasterAddr()
			fmt.Println(addr, err)
			if err != nil {
				return nil, err
			}
			return redis.Dial("tcp", addr, opts.dialOpts...)
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(pool)
	}

	return &mode{
		st,
		pool,
	}
}
