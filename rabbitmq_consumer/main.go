package main

var RabbitMQER *RabbitMQConsumer

func init() {
	RabbitMQER = NewRabbitMQConsumer("amqp://127.0.0.1:5672/")
}

func main() {
	RabbitMQER.Consumer("golang")
}
