package main

import (
	"context"
	"flag"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners/shipping"
	"github.com/medium-stories/go-rabbitmq/order/repository"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var rmqHost = flag.String("rmq_host", "localhost", "RabbitMQ host address")

func main() {
	flag.Parse()

	cred := rmq.NewCredentials()
	cred.Host = *rmqHost
	hub := rmq.NewHub(cred)

	hubCtx, hubCancel := context.WithCancel(context.Background())
	defer hubCancel()

	_, err := hub.Connect(hubCtx)
	if err != nil {
		logrus.Fatal(err)
	}

	orderPaid := shipping.NewOrderPaidListener(hub, repository.NewInMemory())

	event.Listen(hubCtx, orderPaid)

	logrus.Info("listening for messages...")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
