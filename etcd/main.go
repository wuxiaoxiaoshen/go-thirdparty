package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

type ClientEtcd struct {
	config *clientv3.Config
	client *clientv3.Client
}

func NewClientEtcd(host string, port string) *ClientEtcd{
	config := clientv3.Config{
		Endpoints:[]string{host + ":" + port},
		DialTimeout:time.Second*5,
	}
	client,e := clientv3.New(config)
	if e!=nil{
		log.Panic(e)
	}
	return &ClientEtcd{
		config: &config,
		client: client,
	}
}

func main(){
	c := NewClientEtcd("localhost", "2379")
	Example(c)

	// 日本时区 Asia/Tokyo
	t , _:= time.LoadLocation("Asia/Tokyo")

	t.String()
	now := time.Now()
	t1 := now.In(t)
	fmt.Println(now.String())
	fmt.Println(t1.String())
}

func Example(c *ClientEtcd){
	result, e := c.client.Get(context.Background(), "hello")
	if e!=nil{
		log.Println(e)
	}
	//fmt.Println(result)
	for index, i := range result.Kvs{
		fmt.Println(index, string(i.Key), string(i.Value))
	}
	//time.LoadLocation(time.Local)
}