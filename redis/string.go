package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type StringAction struct {
	Key   string
	Value string
	con   *redis.Conn
	mu    sync.Mutex
}

func (S *StringAction) Set() bool {
	con := *S.con
	S.mu.Lock()
	defer S.mu.Unlock()
	_, e := con.Do("SET", S.Key, S.Value)
	if e != nil {
		return false
	}
	return true
}

func (S *StringAction) Get() (string, bool) {
	con := *S.con
	r, e := con.Do("GET", S.Key)
	if e != nil {
		return "-1", false
	}
	result, e := redis.String(r, e)
	return result, true
}

func (S *StringAction) Incr(incr ...int) (int, bool) {
	con := *S.con
	if len(incr) < 1 {
		r, e := con.Do("INCR", S.Key)
		if e != nil {
			log.Println(fmt.Sprintf("redis: incr : %s", e.Error()))
			return -1, false
		}
		rr, e := redis.Int(r, e)
		if e != nil {
			log.Println(fmt.Sprintf("redis: incr : %s", e.Error()))
			return -1, false
		}
		return rr, true
	} else {
		r, e := con.Do("INCRBY", S.Key, incr[0])
		if e != nil {
			log.Println(fmt.Sprintf("redis: incrby : %s", e.Error()))
			return -1, false
		}
		rr, e := redis.Int(r, e)
		if e != nil {
			log.Println(fmt.Sprintf("redis: incrby : %s", e.Error()))
			return -1, false
		}
		return rr, true
	}
}

func (S *StringAction) Decr(incr ...int) (int, bool) {
	con := *S.con
	var ok bool
	if len(incr) >= 1 {
		ok = true
	}
	if ok {
		r, e := con.Do("DECRBY", S.Key, incr[0])
		if e != nil {
			log.Println(fmt.Sprintf("redis: decrby : %s", e.Error()))
			return -1, false
		}
		rr, e := redis.Int(r, e)
		if e != nil {
			log.Println(fmt.Sprintf("redis: decrby : %s", e.Error()))
			return -1, false
		}
		return rr, true
	} else {
		r, e := con.Do("DECR", S.Key)
		if e != nil {
			log.Println(fmt.Sprintf("redis: decr : %s", e.Error()))
			return -1, false
		}
		rr, _ := redis.Int(r, e)
		return rr, true
	}
}

func (S *StringAction) Run(key string, value string, expire int) (interface{}, bool) {
	con := *S.con
	// expire : second
	r, e := con.Do("SET", key, value, "ex", expire, "nx")
	if e != nil {
		log.Println(fmt.Sprintf("redis: set %s ex %d nx", key, expire))
		return nil, false
	}
	result, e := redis.String(r, e)
	if e != nil {
		log.Println(fmt.Sprintf("redis: get value %s %s", key, e.Error()))
		return nil, false
	}
	return result, true

}
