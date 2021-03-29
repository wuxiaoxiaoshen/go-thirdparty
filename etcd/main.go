package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type ClientEtcd struct {
	config *clientv3.Config
	client *clientv3.Client
}

func NewClientEtcd(hosts []string) *ClientEtcd {
	config := clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: time.Second * 5,
	}
	client, e := clientv3.New(config)
	if e != nil {
		log.Panic(e)
	}
	return &ClientEtcd{
		config: &config,
		client: client,
	}
}

func main() {
	c := NewClientEtcd([]string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"})
	Example(c)
}

func Example(c *ClientEtcd) {
	result, e := c.client.Get(context.Background(), "hello")
	if e != nil {
		log.Println(e)
	}
	for index, i := range result.Kvs {
		fmt.Println(index, string(i.Key), string(i.Value))
	}
	result1, _ := c.client.Put(context.Background(), "world", "hello")
	fmt.Println(result1)
}
