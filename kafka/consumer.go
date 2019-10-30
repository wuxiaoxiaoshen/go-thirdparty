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

type ConsumerAction struct {
	c     sarama.ConsumerGroup
	ready chan bool
}

func ini() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
}
func newConsumer(brokerList []string, groupId string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Retry.Backoff = 100 * time.Millisecond
	config.Consumer.Return.Errors = true
	config.Metadata.Retry.Max = 3
	client, err := sarama.NewClient(brokerList, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	consumer, err := sarama.NewConsumerGroupFromClient(groupId, client)
	return consumer
}

func NewConsumer(brokerList []string, groupId string) *ConsumerAction {
	return &ConsumerAction{
		c: newConsumer(brokerList, groupId),
	}
}

func (C ConsumerAction) Close() {
	C.c.Close()
}

func (C *ConsumerAction) Do(topic string, partition int) {
	ctx, cancel := context.WithCancel(context.Background())
	n := ConsumerAction{
		ready: make(chan bool),
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := C.c.Consume(ctx, []string{topic}, &n); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			n.ready = make(chan bool)
		}
	}()
	<-n.ready
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err := C.c.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}

}
func (C *ConsumerAction) Setup(sarama.ConsumerGroupSession) error {
	close(C.ready)
	return nil
}

func (C *ConsumerAction) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (C *ConsumerAction) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
