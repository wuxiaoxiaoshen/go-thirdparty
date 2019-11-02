package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaAction struct {
	DataSyncProducer  sarama.SyncProducer
	DataAsyncProducer sarama.AsyncProducer
}

var TOPIC string

func newDataSyncProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 5                    // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer1:", err)
	}
	return producer

}

func newDataAsyncProducer(brokerList []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stdout, "[KAFKA] ", log.LstdFlags)
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer2:", err)
	}
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()
	return producer
}

func NewKafkaAction(brokerList []string) *KafkaAction {
	return &KafkaAction{
		DataSyncProducer:  newDataSyncProducer(brokerList),
		DataAsyncProducer: newDataAsyncProducer(brokerList),
	}
}

func (K *KafkaAction) Do(v interface{}) {
	message := v.(SendMessage)
	partition, offset, err := K.DataSyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: TOPIC,
		Value: &message,
	})
	if err != nil {
		log.Println(err)
		return
	}
	value := map[string]string{
		"method": message.Method,
		"url":    message.URL,
		"value":  message.Value,
		"date":   message.Date,
	}
	fmt.Println(fmt.Sprintf("/%d/%d/%+v", partition, offset, value))
}

func (K *KafkaAction) String() string {
	return ""
}

func (K *KafkaAction) Run(v interface{}) {
	message := v.(SendMessage)
	K.DataAsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: TOPIC,
		Value: &message,
	}
}

func (K *KafkaAction) Close() {
	e := K.DataAsyncProducer.Close()
	if e != nil {
		log.Println("Failed to shut down data sync collector cleanly", e)
	}
	e = K.DataAsyncProducer.Close()
	if e != nil {
		log.Println("Failed to shut down data async collector cleanly", e)
	}
}
