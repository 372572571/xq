package redisx

import "github.com/gomodule/redigo/redis"

type (
	Conn interface {
		// Close closes the connection.
		Close() error

		// Err returns a non-nil value when the connection is not usable.
		Err() error

		// Do sends a command to the server and returns the received reply.
		// This function will use the timeout which was set when the connection is created
		Do(commandName string, args ...interface{}) (reply interface{}, err error)

		// Send writes the command to the client's output buffer.
		Send(commandName string, args ...interface{}) error

		// Flush flushes the output buffer to the Redis server.
		Flush() error

		// Receive receives a single reply from the Redis server
		Receive() (reply interface{}, err error)
	}
	Mode interface {
		GetConn() redis.Conn
		NewConn() (redis.Conn, error)
		GetPool() *redis.Pool
	}

	Redis struct {
		mode Mode
		impl *redis.Pool
		Values
	}
)

func (r *Redis) Impl() *redis.Pool {
	return r.impl
}

func (r *Redis) Do(next func(conn Conn)) {
	var p = r.impl.Get()
	defer p.Close()
	next(p)
}

func (r *Redis) DoEx(next func(conn Conn) error) error {
	var p = r.impl.Get()
	defer p.Close()
	return next(p)
}
