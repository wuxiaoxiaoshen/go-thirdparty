package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type BrokerAction struct {
	Addr   string
	broker *sarama.Broker
}

func NewBrokerAction(addr string) *BrokerAction {
	broker := sarama.NewBroker(addr)
	return &BrokerAction{
		Addr:   addr,
		broker: broker,
	}
}

func (B *BrokerAction) GetMetaMessage(topic string) interface{} {
	req := sarama.MetadataRequest{
		Topics:                 []string{topic},
		AllowAutoTopicCreation: false,
	}
	B.broker.Open(nil)
	response, e := B.broker.GetMetadata(&req)
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
	B.broker.Open(nil)
	r, e := B.broker.ListGroups(&g)
	if e != nil {
		log.Println(fmt.Sprintf("[saram List group: %s]", e.Error()))
		return e
	}
	return r

}
func (B *BrokerAction) Close() {
	B.broker.Close()
}
