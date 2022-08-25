package main

import (
	"context"
	"flag"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	listeners "github.com/medium-stories/go-rabbitmq/event/listeners/order"
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

	orderCreated := listeners.NewOrderCreatedListener(hub)
	orderPaid := listeners.NewOrderPaidListener(hub, repository.NewSqlite("woohoo_orders"))
	orderShipped := listeners.NewOrderShippedListener(hub, repository.NewSqlite("woohoo_orders"))

	event.Listen(hubCtx, orderCreated, orderPaid, orderShipped)

	logrus.Info("listening for messages...")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
