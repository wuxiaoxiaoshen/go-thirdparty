package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaAction struct {
	DataSyncProducer  sarama.SyncProducer
	DataAsyncProducer sarama.AsyncProducer
}

func newDataSyncProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 5                    // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	return producer

}

func newDataAsyncProducer(brokerList []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
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
		Topic: "data_sync",
		Key:   sarama.StringEncoder("data_sync_key"),
		Value: &message,
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("/%d/%d", partition, offset))
}

func (K *KafkaAction) String() string {
	return ""
}

func (K *KafkaAction) Run(v interface{}) {
	message := v.(SendMessage)
	K.DataAsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: "data_async",
		Key:   sarama.StringEncoder("data_async"),
		Value: message,
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
