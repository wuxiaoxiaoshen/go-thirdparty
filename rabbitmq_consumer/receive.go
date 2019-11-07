package main

import (
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn *amqp.Connection
}

func NewRabbitMQConsumer(addr string) *RabbitMQConsumer {
	conn, e := amqp.Dial(addr)
	if e != nil {
		log.Println(e)
		panic(e)
		return nil
	}
	return &RabbitMQConsumer{conn: conn}
}

func (R *RabbitMQConsumer) queueDeclare(topic string) *amqp.Queue {
	ch, e := R.conn.Channel()
	if e != nil {
		log.Println(e)
		return nil
	}
	q, e := ch.QueueDeclare(topic, false, false, false, false, nil)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &q
}

func (R *RabbitMQConsumer) Consumer(topic string) bool {
	ch, e := R.conn.Channel()
	if e != nil {
		log.Println(e)
		return false
	}
	q := R.queueDeclare(topic)
	var wg sync.WaitGroup
	msgs, e := ch.Consume(q.Name, "", true, false, false, false, nil)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	wg.Wait()
	return true
}
