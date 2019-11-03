package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

var ConsumerAction *NsqConsumerAction

func init() {
	ConsumerAction = NewNsqConsumerAction("topic-nsq")
}

var APP *iris.Application

func init() {
	APP = iris.New()
	APP.Logger().SetLevel("debug")
	APP.Use(recover.New())
}

func Register(i iris.Party) {
	i.Get("/consumer", func(context iris.Context) {
		context.JSON(iris.Map{
			"data": "consumer consumer",
		})
	})
}

func main() {
	APP.PartyFunc("/v1/api/nsq", Register)
	go func() {
		ConsumerAction.Run("0.0.0.0:32772")
	}()
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
			APP.Shutdown(ctx)
		}
	}()
	APP.Run(iris.Addr(":8082"), iris.WithoutServerError(iris.ErrServerClosed))

}
