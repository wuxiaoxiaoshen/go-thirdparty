package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type KeyAction struct {
	con *redis.Conn
}

func NewKeyAction(con redis.Conn) *KeyAction {
	return &KeyAction{con: &con}
}

func (k *KeyAction) Exists(key string) bool {
	c := *k.con
	ok, e := redis.Bool(c.Do("EXISTS", key))
	if e != nil {
		log.Println(fmt.Sprintf("redis: exists %s", key))
		return false
	}
	return ok

}
func (k *KeyAction) DBSize() int {
	c := *k.con
	size, e := redis.Int(c.Do("DBSIZE"))
	if e != nil {
		log.Println(fmt.Sprintf("redis: dbsize %d", size))
		return 0
	}
	return size
}
func (k *KeyAction) FlushAll() bool {
	c := *k.con
	_, e := c.Do("FLUSHALL")
	if e != nil {
		log.Println(fmt.Sprintf("redis: flushall %s", e.Error()))
		return false
	}
	return true
}
func (k *KeyAction) Type(key string) string {
	c := *k.con
	t, e := redis.String(c.Do("TYPE", key))
	if e != nil {
		log.Println(fmt.Sprintf("redis: type %s %s", key, e.Error()))
		return "None"
	}
	return t
}

func (k *KeyAction) Del(key string) bool {
	c := *k.con
	r, e := redis.Bool(c.Do("DEL", key))
	if e != nil {
		log.Println(fmt.Sprintf("redis: del %s %s", key, e.Error()))
		return false
	}
	return r
}
