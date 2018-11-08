package redis

import (
	"github.com/gomodule/redigo/redis"
)

var redisPool = &redis.Pool{
	MaxIdle:   10,
	MaxActive: 100,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			panic(err.Error())
		}
		return c, err
	},
}

func perform(fn func(c redis.Conn) (interface{}, error)) (interface{}, error) {
	pool := redisPool
	conn := pool.Get()
	defer conn.Close()
	reply, err := fn(conn)
	return reply, err
}

func SET(k string, v string, ttl int64) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("SET", k, v, "EX", ttl, "NX")
		return reply, err
	}
	reply, err := perform(fn)
	return reply, err
}

func GET(k string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("GET", k)
		return reply, err
	}
	reply, err := perform(fn)
	return reply, err
}

func DEL(k string) (interface{}, error) {
	fn := func(c redis.Conn) (interface{}, error) {
		reply, err := c.Do("DEL", k)
		return reply, err
	}
	reply, err := perform(fn)
	return reply, err
}