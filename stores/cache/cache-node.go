package cache

import (
	"context"
	"errors"
	"time"

	"github.com/372572571/xq/jsonx"
	"github.com/372572571/xq/logx"
	"github.com/372572571/xq/stores/redisx"
	"github.com/372572571/xq/syncx"
)

const (
	notFoundPlaceholder = "*"
	// // make the expiry unstable to avoid lots of cached items expire at the same time
	// // make the unstable expiry to be [0.95, 1.05] * seconds
	// expiryDeviation = 0.05
)

// indicates there is no such value associate with the key
var errPlaceholder = errors.New("placeholder")

type (
	cacheNode struct {
		rds            *redisx.Redis
		expiry         time.Duration // default expiry time
		notFoundExpiry time.Duration // 站位锁过期时间
		barrier        syncx.SingleCall
		// r              *rand.Rand
		// lock           *sync.Mutex
		errNotFound error
	}
)

func NewNode(rds *redisx.Redis, barrier syncx.SingleCall,
	errNotFound error, opts ...Option) interface{} {
	o := newOptions(opts...)
	return cacheNode{
		rds:            rds,
		expiry:         o.Expiry,
		notFoundExpiry: o.NotFoundExpiry,
		barrier:        barrier,
		// r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		// lock:           new(sync.Mutex),
		errNotFound: errNotFound,
	}
}

func (c cacheNode) IsNotFound(err error) bool {
	return errors.Is(err, c.errNotFound)
}

func (c cacheNode) doGetCache(ctx context.Context,
	key string, v interface{}) error {
	if err := c.rds.DoEx(func(conn redisx.Conn) error {
		value, err := c.rds.String(conn.Do("GET", key))
		if err != nil {
			return err
		}
		if len(value) == 0 {
			return c.errNotFound
		}
		if value == notFoundPlaceholder {
			return errPlaceholder
		}
		err = c.processCache(ctx, key, value, v)
		return err
	}); err != nil {
		return err
	}
	return nil
}

func (c cacheNode) Take(val interface{}, key string, query func(val interface{}) error) error {
	return c.TakeCtx(context.Background(), val, key, query)
}

func (c cacheNode) TakeCtx(ctx context.Context, val interface{}, key string, query func(val interface{}) error) error {
	return c.doTake(ctx, val, key, query, func(v interface{}) error {
		return c.SetCtx(ctx, key, v)
	})
}

func (c cacheNode) TakeWithExpire(val interface{}, key string, query func(val interface{}, expire time.Duration) error) error {
	return c.TakeWithExpireCtx(context.Background(), val, key, query)
}

func (c cacheNode) TakeWithExpireCtx(ctx context.Context, val interface{}, key string, query func(val interface{}, expire time.Duration) error) error {
	cacheval := func(v interface{}) error {
		return c.SetWithExpireCtx(ctx, key, v, c.expiry)
	}
	call_query := func(v interface{}) error {
		return query(v, c.expiry)
	}
	return c.doTake(ctx, val, key, call_query, cacheval)
}

func (c cacheNode) doTake(ctx context.Context, v interface{}, key string,
	query func(v interface{}) error, cacheVal func(v interface{}) error) error {

	val, fresh, err := c.barrier.DoEx(key, func() (interface{}, error) {

		if err := c.doGetCache(ctx, key, v); err != nil {
			if err == errPlaceholder {
				return nil, c.errNotFound
			} else if err != c.errNotFound {
				return nil, err
			}

			if err = query(v); err == c.errNotFound {

				if err = c.setCacheWithNotFound(ctx, key); err != nil {
					logx.Error(err)
				}

				return nil, c.errNotFound

			} else if err != nil {
				return nil, err
			}

			// call the callback setting cache value
			if err = cacheVal(v); err != nil {
				logx.Error(err)
			}
		}

		return jsonx.Marshal(v)
	})

	if err != nil {
		return err
	}

	// not yourself get but can use
	if fresh {
		return nil
	}

	return jsonx.Unmarshal(val.([]byte), v)
}

// delete the selected cache
func (c cacheNode) Del(keys ...string) error {
	return c.DelCtx(context.Background(), keys...)
}

func (c cacheNode) DelCtx(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	c.rds.Do(func(c redisx.Conn) {
		for _, key := range keys {
			if _, err := c.Do("DEL", key); err != nil {
				logx.Errorf("cache delete key:%s err:%s", key, err.Error())
			}
		}
		return
	})
	return nil
}

func (c cacheNode) Get(key string, val interface{}) error {
	return c.GetCtx(context.Background(), key, val)
}

func (c cacheNode) GetCtx(ctx context.Context, key string, val interface{}) error {
	err := c.doGetCache(ctx, key, val)
	if err == errPlaceholder {
		return c.errNotFound
	}

	return err
}

func (c cacheNode) processCache(ctx context.Context,
	key, data string, v interface{}) error {
	err := jsonx.Unmarshal([]byte(data), v)
	if err == nil {
		return nil
	}

	// 到这错误说明内容格式不正常
	// 删除异常数据
	logx.Errorf("cache value not json key:%s value:%s err:%s",
		key, data, err.Error())
	// 尝试删除错误缓存
	c.DelCtx(ctx, key)
	return err
}

// setting occupy a seat
func (c cacheNode) setCacheWithNotFound(ctx context.Context, key string) error {
	if err := c.rds.DoEx(func(conn redisx.Conn) error {
		_, err := conn.Do("SET", key, notFoundPlaceholder,
			"EX", c.notFoundExpiry.Seconds())

		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil
	}
	return nil
}

func (c cacheNode) Set(key string,
	val interface{}) error {

	return c.SetWithExpireCtx(context.Background(),
		key, val, time.Duration(c.expiry.Seconds()))
}

func (c cacheNode) SetCtx(ctx context.Context, key string,
	val interface{}) error {

	return c.SetWithExpireCtx(ctx, key, val,
		time.Duration(c.expiry.Seconds()))
}

func (c cacheNode) SetWithExpire(key string,
	val interface{}, expire time.Duration) error {

	return c.SetWithExpireCtx(context.Background(), key, val, expire)
}

func (c cacheNode) SetWithExpireCtx(ctx context.Context, key string,
	val interface{}, expire time.Duration) error {

	if err := c.rds.DoEx(func(conn redisx.Conn) error {
		data, err := jsonx.Marshal(val)
		if err != nil {
			return err
		}
		_, err = conn.Do("SET", key, data, "EX", c.notFoundExpiry.Seconds())
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
