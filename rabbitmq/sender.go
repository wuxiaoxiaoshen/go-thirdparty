package main

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQAction struct {
	conn *amqp.Connection
}

func NewRabbitMQAction(addr string) *RabbitMQAction {
	conn, e := amqp.Dial(addr)
	if e != nil {
		log.Println(e)
		panic(e)
		return nil
	}
	return &RabbitMQAction{conn: conn}
}

func (R *RabbitMQAction) Close() {
	defer R.conn.Close()
}
func (R *RabbitMQAction) queueDeclare(name string) *amqp.Queue {
	ch, e := R.conn.Channel()
	if e != nil {
		log.Println(e)
		return nil
	}
	q, e := ch.QueueDeclare(name, false, false, false, false, nil)
	if e != nil {
		log.Println(e)
		return nil
	}
	return &q
}

func (R *RabbitMQAction) Publish(name string, body string) bool {
	ch, e := R.conn.Channel()
	if e != nil {
		log.Println(e)
		return false
	}
	q := R.queueDeclare(name)
	e = ch.Publish("", q.Name, false, false,
		amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if e != nil {
		log.Println(e)
		return false
	}
	return true
}
