package main

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type AdminAction struct {
	c sarama.ClusterAdmin
}

func NewAdminAction(addr []string) *AdminAction {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	sarama.Logger = log.New(os.Stdout, "[sarama ] ", log.LstdFlags)
	admin, e := sarama.NewClusterAdmin(addr, config)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &AdminAction{c: admin}

}

func (A *AdminAction) GetTopic() interface{} {
	detail, e := A.c.ListTopics()
	if e != nil {
		log.Println(e)
		return nil
	}
	return detail
}
func (A *AdminAction) GetGroups() interface{} {
	detail, e := A.c.ListConsumerGroups()
	if e != nil {
		log.Println(e)
		return nil
	}
	return detail
}
func (A *AdminAction) CreateTopic(topic string, partition int32, factor int16) bool {
	detail := sarama.TopicDetail{
		NumPartitions:     partition,
		ReplicationFactor: factor,
	}
	e := A.c.CreateTopic(topic, &detail, false)
	if e != nil {
		log.Println(e)
		return false
	}
	return true
}

func (A *AdminAction) DescribeTopic(topic []string) interface{} {
	detail, e := A.c.DescribeTopics(topic)
	if e != nil {
		log.Println(e)
		return nil
	}
	return detail
}

func (A *AdminAction) DeleteTopic(topic string) bool {
	e := A.c.DeleteTopic(topic)
	if e != nil {
		log.Println(e)
		return false
	}
	return true

}
