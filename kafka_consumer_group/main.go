package main

import (
	"context"
	"sync"
)

var KafkaConsumerGroup *KafkaConsumerGroupAction

func init() {
	KafkaConsumerGroup = NewKafkaConsumerGroupAction([]string{"localhost:9092"}, "Siren-Production-2")
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithCancel(context.Background())
	KafkaConsumerGroup.Consume([]string{"frequent_customer_dev", "store_frequent_customer_dev"}, *wg, ctx)
}
