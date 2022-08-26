package order

import (
	"context"
	"encoding/json"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners"
	"github.com/sirupsen/logrus"
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
	consumer := listeners.StartConsumer(ctx, ev.hub, "order", event.OrderCreated)
	ev.handleMessages(ctx, consumer, event.OrderCreated)
}

func (ev *orderCreated) handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			var identifier string
			json.Unmarshal(msg, &identifier)

			logrus.Infof("[%s] - %s", name, identifier)
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
