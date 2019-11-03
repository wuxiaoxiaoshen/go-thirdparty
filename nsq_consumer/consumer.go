package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nsqio/go-nsq"
)

type NsqConsumerAction struct {
	consumer *nsq.Consumer
}

func NewNsqConsumerAction(topic string) *NsqConsumerAction {
	config := nsq.NewConfig()
	consumer, e := nsq.NewConsumer(topic, "default", config)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &NsqConsumerAction{consumer: consumer}
}

func (N *NsqConsumerAction) Do(addr string) {
	N.consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("nsqd-1", message.Timestamp, message.NSQDAddress, string(message.Body))
		return nil
	}))
	//服务发现
	err := N.consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Println(err)
	}
}
func (N *NsqConsumerAction) Run(addr string) {

	N.consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("nsqd-2", message.Timestamp, message.NSQDAddress, string(message.Body))
		return nil
	}))
	log.Println("ddd")
	// 直连
	err := N.consumer.ConnectToNSQD(addr)
	if err != nil {
		log.Fatal(err)
	}
	stats := N.consumer.Stats()
	if stats.Connections == 0 {
		panic("stats report 0 connections (should be > 0)")
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("server is running....")
	<-stop
}
