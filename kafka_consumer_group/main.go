package main

import (
	"context"
	"sync"
)

var KafkaConsumerGroup *KafkaConsumerGroupAction

func init() {
	KafkaConsumerGroup = NewKafkaConsumerGroupAction([]string{"127.0.0.1:9092"}, "customer-group-2")
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithCancel(context.Background())
	KafkaConsumerGroup.Consume([]string{"topic-python", "topic-golang"}, *wg, ctx)
}
