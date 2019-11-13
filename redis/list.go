package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type ListAction struct {
	con *redis.Conn
}

func NewListAction(con *redis.Conn) *ListAction {
	return &ListAction{con: con}
}

func (L *ListAction) LPush(key string, fields ...interface{}) bool {
	c := *L.con
	for _, i := range fields {
		r, e := c.Do("LPUSH", key, i)
		_, e = redis.Bool(r, e)
		if e != nil {
			log.Println(fmt.Sprintf("redis: lpush %s %v", key, fields))
			return false
		}
	}

	return true
}

func (L *ListAction) RPush(key string, fields ...interface{}) bool {
	c := *L.con
	for _, i := range fields {
		r, e := c.Do("RPUSH", key, i)
		_, e = redis.Bool(r, e)
		if e != nil {
			log.Println(fmt.Sprintf("redis: rpush %s %v", key, fields))
			return false
		}
	}

	return true
}

func (L *ListAction) LPop(key string) ([]interface{}, bool) {
	c := *L.con
	r, e := c.Do("LPOP", key)
	if e != nil {
		log.Println(fmt.Sprintf("redis: rpush %s %v", key, e.Error()))
		return nil, false
	}
	v, e := redis.Values(r, e)
	return v, true

}

//func (L *ListAction) RPop(key string) (interface{}, bool)                     {}
//func (L *ListAction) Len(key string) int                                      {}
//func (L *ListAction) LIndex(key string, index int) (interface{}, bool)        {}
func (L *ListAction) LRange(key string, start, end int) (string, bool) {
	c := *L.con
	r, e := c.Do("LRANGE", key, start, end)
	v, e := redis.Values(r, e)
	if e != nil {
		log.Println(fmt.Sprintf("redis: lrange %s %v", key, e.Error()))
		return "[nil]", false
	}
	var b bytes.Buffer
	b.WriteString("[")
	for index, i := range v {
		if index != len(v)-1 {
			b.WriteString(fmt.Sprintf("%s,", i))
		} else {
			b.WriteString(fmt.Sprintf("%s", i))
		}
	}
	b.WriteString("]")

	return b.String(), true
}

//func (L *ListAction) LTrim(key string, start, end int) ([]interface{}, bool)  {}
