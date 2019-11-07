package main

import (
	"context"
	"sync"
)

var KafkaConsumerGroup *KafkaConsumerGroupAction

func init() {
	KafkaConsumerGroup = NewKafkaConsumerGroupAction([]string{"127.0.0.1:19092"}, "customer-group-1")
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithCancel(context.Background())
	KafkaConsumerGroup.Consume([]string{"topic-python", "topic-golang", "topic-js"}, *wg, ctx)
}
