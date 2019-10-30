package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

var Server *KafkaAction

func init() {
	// broker: 代表的就是 kafka 主机
	Server = NewKafkaAction([]string{"127.0.0.1:9092"})
}

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	return app
}
func party(c iris.Party) {
	c.Post("/kafka/producer", func(context iris.Context) {
		var message SendMessage
		if err := context.ReadJSON(&message); err != nil {
			log.Println(err)
			return
		}
		Server.Do(message)
		Server.Run(message)
		context.JSON(iris.Map{
			"data": message,
		})

	})
	c.Get("/kafka/broker/{topic:string}", func(i iris.Context) {
		topic := i.Params().GetString("topic")
		b := NewBrokerAction("39.107.123.21:9092")
		r := b.GetMetaMessage(topic)
		b.Close()
		i.JSON(iris.Map{
			"data": r,
		})
	})
	c.Get("/kafka/broker/list_group", func(i iris.Context) {
		b := NewBrokerAction("39.107.123.21:9092")
		r := b.GetListGroup()
		b.Close()
		i.JSON(iris.Map{
			"data": r,
		})
	})
	c.Get("/kafka/consumer", func(i iris.Context) {
		c := NewConsumer([]string{"127.0.0.1:9092"}, "Group_test")
		c.Do("data_sync", 0)
		i.JSON(iris.Map{
			"data": "success",
		})

	})
}
func main() {
	app := newApp()
	app.PartyFunc("/v1/api", party)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			os.Interrupt,
			syscall.SIGINT,
			os.Kill,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
