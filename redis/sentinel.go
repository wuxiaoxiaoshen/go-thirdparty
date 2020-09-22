package main

import "github.com/go-redis/redis/v7"

type SentinelAction struct {
	Client *redis.Client
}

func NewSentinelAction(masterName string, address []string, password string) *SentinelAction{
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: address,
		Password: password,
	})

	return &SentinelAction{
		Client: rdb,
	}
}
