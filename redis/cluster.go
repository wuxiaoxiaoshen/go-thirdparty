package main

import "github.com/go-redis/redis/v7"

type ClusterAction struct {
	Client *redis.ClusterClient
}

func NewClusterAction() *ClusterAction{
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	return &ClusterAction{
		Client: rdb,
	}
}
