package main

import "github.com/gomodule/redigo/redis"

type ListAction struct {
	con *redis.Conn
}

func NewListAction(con *redis.Conn) *ListAction {
	return &ListAction{con: con}
}

func (L *ListAction) LPush(key string, fields ...string) ([]interface{}, bool) {

}

func (L *ListAction) RPush(key string, fields ...string) ([]interface{}, bool) {}

func (L *ListAction) LPop(key string) (interface{}, bool)                     {}
func (L *ListAction) RPop(key string) (interface{}, bool)                     {}
func (L *ListAction) Len(key string) int                                      {}
func (L *ListAction) LIndex(key string, index int) (interface{}, bool)        {}
func (L *ListAction) LRange(key string, start, end int) ([]interface{}, bool) {}
func (L *ListAction) LTrim(key string, start, end int) ([]interface{}, bool)  {}
