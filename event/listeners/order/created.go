package order

import (
	"context"
	"fmt"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners"
	"github.com/sirupsen/logrus"
	"time"
)

type orderCreated struct {
	hub *rmq.Hub
}

func NewOrderCreatedListener(hub *rmq.Hub) *orderCreated {
	return &orderCreated{
		hub: hub,
	}
}

func (ev *orderCreated) Listen(ctx context.Context) {
	consumer := listeners.StartConsumer(ctx, ev.hub, event.OrderCreated)
	go ev.handleMessages(ctx, consumer, fmt.Sprintf("order[%s]", event.OrderCreated))
}

func (ev *orderCreated) handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			logrus.Infof("[%s] %s - %s", time.Now().UTC(), name, msg)
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
