package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type PoolAction struct {
	Pool redis.Pool
}

func NewPoolAction(addr string) *PoolAction {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", addr)
		},
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &PoolAction{Pool: pool}
}

func (P *PoolAction) Close() {
	defer P.Pool.Close()
}

func (P *PoolAction) Get() redis.Conn {
	return P.Pool.Get()
}

