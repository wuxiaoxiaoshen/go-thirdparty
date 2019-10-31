package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type BrokerAction struct {
	Addr   string
	broker []*sarama.Broker
	topic  []string
}

func NewBrokerAction(addr string) *BrokerAction {
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	client, err := sarama.NewClient([]string{addr}, config)
	if err != nil {
		log.Println(err)
		return nil
	}
	topic, _ := client.Topics()
	broker := client.Brokers()
	return &BrokerAction{
		Addr:   addr,
		broker: broker,
		topic:  topic,
	}
}

func (B *BrokerAction) GetMetaMessage(topic string) interface{} {
	req := sarama.MetadataRequest{
		Topics:                 []string{topic},
		AllowAutoTopicCreation: false,
	}
	B.broker[0].Open(nil)
	response, e := B.broker[0].GetMetadata(&req)
	if e != nil {
		log.Println(fmt.Sprintf("[saram broker get meta message]: %s", e.Error()))
		return e
	}
	return response
}

func (B *BrokerAction) HearBeat(topic string) interface{} {

	return nil
}
func (B *BrokerAction) GetListGroup() interface{} {
	g := sarama.ListGroupsRequest{}
	B.broker[0].Open(nil)
	r, e := B.broker[0].ListGroups(&g)
	if e != nil {
		log.Println(fmt.Sprintf("[saram List group: %s]", e.Error()))
		return e
	}
	return r

}

func (B *BrokerAction) CreatTopic(topic string, partitions int32, replicationFactor int16) interface{} {
	detail := sarama.TopicDetail{
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}
	topicDetail := make(map[string]*sarama.TopicDetail)
	topicDetail[topic] = &detail
	topicRequest := sarama.CreateTopicsRequest{
		TopicDetails: topicDetail,
	}
	r, e := B.broker[0].CreateTopics(&topicRequest)
	if e != nil {
		log.Println(e)
		return e
	}
	return r
}

func (B *BrokerAction) DeleteTopic(topic string) bool {
	detail := sarama.DeleteTopicsRequest{
		Topics: []string{topic},
	}
	_, e := B.broker[0].DeleteTopics(&detail)
	if e != nil {
		log.Println(e)
		return false
	}
	return true
}

func (B *BrokerAction) GetTopics() []string {
	return B.topic
}

func (B *BrokerAction) Close() {
	for _, i := range B.broker {
		i.Close()
	}
}
