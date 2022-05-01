package sentinel

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type options struct {
	addrs            []string                 // # 哨兵集群地址 #
	masterName       string                   // # 主节点名称 #
	poolOpts         []func(pool *redis.Pool) // # 连接池配置函数 #
	dialOpts         []redis.DialOption       // # redis拨号,配置函数 #
	sentinelDialOpts []redis.DialOption       // # 哨兵拨号,配置函数 #
}

type OptFunc func(opts *options)

func Addrs(value []string) OptFunc {
	return func(opts *options) {
		fmt.Println(opts.addrs)
		opts.addrs = value
	}
}

func MasterName(value string) OptFunc {
	return func(opts *options) {
		fmt.Println(opts.masterName)
		opts.masterName = value
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

func DialSentinelOpts(value ...redis.DialOption) OptFunc {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.sentinelDialOpts = append(opts.sentinelDialOpts, dialOpt)
		}
	}
}
