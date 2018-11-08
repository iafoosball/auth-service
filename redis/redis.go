package redis

import (
	"github.com/gomodule/redigo/redis"
)

var redisPool = &redis.Pool {
	MaxIdle:10,
	MaxActive: 100,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			panic(err.Error())
		}
		return c, err
	},
}

func Perform(command string, args ...interface{}) (interface{}, error) {
	pool := redisPool
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do(command, args)
	if err != nil {
		return nil, err
	}
	return reply, err
}



