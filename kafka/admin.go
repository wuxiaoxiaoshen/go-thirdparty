package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

type AdminAction struct {
	c sarama.ClusterAdmin
}

func NewAdminAction(addr []string) *AdminAction {
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Net.SASL.Version = sarama.SASLHandshakeV1
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 10 * time.Second
	config.Metadata.Retry.Max = 3
	config.Net.ReadTimeout = 30 * time.Second
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	sarama.Logger = log.New(os.Stdout, "[sarama ] ", log.LstdFlags)
	admin, e := sarama.NewClusterAdmin(addr, config)
	if e != nil {
		log.Println(e, "ERROR")
		return nil
	}
	fmt.Println(admin.ListTopics())
	return &AdminAction{c: admin}

}

func (A *AdminAction) GetTopic() interface{} {
	detail, e := A.c.ListTopics()
	if e != nil {
		log.Println("e", e)
		return nil
	}
	return detail
}
func (A *AdminAction) GetGroups(groupID []string) interface{} {
	detail, e := A.c.DescribeConsumerGroups(groupID)
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
