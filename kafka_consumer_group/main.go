package main

import (
	"context"
	"sync"
)

var KafkaConsumerGroup *KafkaConsumerGroupAction

func init() {
	KafkaConsumerGroup = NewKafkaConsumerGroupAction([]string{"47.93.81.180:9092"}, "Siren-Production-1")
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithCancel(context.Background())
	KafkaConsumerGroup.Consume([]string{"frequent_customer_production", "store_frequent_customer_production"}, *wg, ctx)
}
