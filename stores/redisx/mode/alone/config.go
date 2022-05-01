package alone

import (
	"github.com/gomodule/redigo/redis"
)

type options struct {
	addr     string                   // # 地址 #
	poolOpts []func(pool *redis.Pool) // # 连接池配置函数 #
	dialOpts []redis.DialOption       // # redis拨号,配置函数 #
}

type OptFunc func(opts *options)

func Addrs(value string) OptFunc {
	return func(opts *options) {
		opts.addr = value
	}
}

func PoolOpts(value ...func(pool *redis.Pool)) OptFunc {
	return func(opts *options) {
		for _, poolOpt := range value {
			opts.poolOpts = append(opts.poolOpts, poolOpt)
		}
	}
}

func DialOpts(value ...redis.DialOption) OptFunc {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.dialOpts = append(opts.dialOpts, dialOpt)
		}
	}
}
