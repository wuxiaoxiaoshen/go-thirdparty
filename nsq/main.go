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

var NsqProducerAction *ProducerAction

func init() {
	NsqProducerAction = NewProducerAction("0.0.0.0:32772")
}

type APP struct {
	app *iris.Application
}

func NewAPP() *APP {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Logger().SetLevel("debug")
	return &APP{app: app}
}

func (A *APP) Party() {
	A.app.PartyFunc("/v1/api/nsq", A.Register)
}

func (A *APP) Register(i iris.Party) {
	i.Post("/producer/{topic:string}", func(i iris.Context) {
		topic := i.Params().GetString("topic")
		var msg Message
		if e := i.ReadJSON(&msg); e != nil {
			log.Println(e)
			return
		}
		body, e := msg.Encoder()
		if e != nil {
			log.Println(e)
			return
		}
		NsqProducerAction.Do(topic, body)
		i.JSON(iris.Map{
			"data": "success",
		})

	})
}
func main() {
	app := NewAPP()
	app.Party()
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
			app.app.Shutdown(ctx)
		}
	}()
	app.app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}
