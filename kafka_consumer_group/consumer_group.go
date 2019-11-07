package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaConsumerGroupAction struct {
	group sarama.ConsumerGroup
}

func NewKafkaConsumerGroupAction(brokers []string, groupId string) *KafkaConsumerGroupAction {
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stdout, "[consumer_group]", log.Lshortfile)
	// 重平衡策略
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 6 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 2 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V2_0_0_0
	consumerGroup, e := sarama.NewConsumerGroup(brokers, groupId, config)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &KafkaConsumerGroupAction{group: consumerGroup}

}

func (K *KafkaConsumerGroupAction) Consume(topics []string, wg sync.WaitGroup, ctx context.Context) {
	var consumer = KafkaConsumerGroupHandler{ready: make(chan bool)}
	go func() {
		defer wg.Done()
		for {
			if err := K.group.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	log.Println("Sarama consumer up and running!...")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	wg.Wait()
	if err := K.group.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

type KafkaConsumerGroupHandler struct {
	ready chan bool
}

func (K *KafkaConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (K *KafkaConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (K *KafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, partions = %d, offset = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}

	return nil
}
