package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type ProducerAction struct {
	addr     string
	producer *nsq.Producer
}

func NewProducerAction(addr string) *ProducerAction {
	config := nsq.NewConfig()
	producer, e := nsq.NewProducer(addr, config)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &ProducerAction{
		addr:     addr,
		producer: producer,
	}
}

func (P *ProducerAction) Do(topic string, body []byte) {
	if e := P.producer.Publish(topic, body); e != nil {
		log.Println(e)
		return
	}
}

func (P *ProducerAction) String() string {
	return ""
}
