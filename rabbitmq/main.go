package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var Application *iris.Application
var RabbitMQER *RabbitMQAction

func init() {
	Application = iris.New()
	Application.Use(recover.New())
	Application.Use(logger.New())
	Application.Logger().SetLevel("debug")
	RabbitMQER = NewRabbitMQAction("amqp://127.0.0.1:5672/")
}

func Party() {
	Application.PartyFunc("/v1/api/amqp", Register)
}
func Register(c iris.Party) {
	c.Post("/sender/{topic:string}", func(i iris.Context) {
		topic := i.Params().GetString("topic")
		var params struct {
			Body string
		}
		e := i.ReadJSON(&params)
		if e != nil {
			log.Println(e)
			return
		}
		ok := RabbitMQER.Publish(topic, params.Body)
		i.JSON(iris.Map{
			"data": ok,
		})
	})
}

func main() {
	Party()
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
			Application.Shutdown(ctx)
		}
	}()
	Application.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
